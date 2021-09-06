package service

import (
	"context"
	ghttp "net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"

	pbuser "thooh/api/common/v1/user"
	pberror "thooh/api/error"
	pb "thooh/api/wechat/v1/base"
	"thooh/internal/biz"
	"thooh/internal/conf"
)

type WeChatService struct {
	pb.UnimplementedWeChatServer
	officialAccount *officialaccount.OfficialAccount
	userUsecase     *biz.UserUsecase
	logger          *log.Helper
	wechatUsecase   *biz.WechatUsecase
}

func NewWeChatService(
	logger log.Logger,
	c *conf.Wechat,
	userUsecase *biz.UserUsecase,
	wechatUsecase *biz.WechatUsecase,
) *WeChatService {
	wc := wechat.NewWechat()
	// 这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          c.AppId,
		AppSecret:      c.AppSecret,
		Token:          c.Token,
		EncodingAESKey: c.EncodingAesKey,
		Cache:          memory,
	}
	return &WeChatService{
		officialAccount: wc.GetOfficialAccount(cfg),
		userUsecase:     userUsecase,
		logger:          log.NewHelper(logger),
		wechatUsecase:   wechatUsecase,
	}
}

func (s *WeChatService) QRCode(ctx context.Context, req *pb.QRCodeReq) (*pb.QRCodeResp, error) {
	resp := new(pb.QRCodeResp)
	ticket, err := s.officialAccount.GetBasic().GetQRTicket(&basic.Request{
		ExpireSeconds: int64(req.GetExpireSeconds()),
		ActionName:    pb.QRCodeReq_ActionName_name[int32(req.GetActionName())],
		ActionInfo: struct {
			Scene struct {
				SceneStr string `json:"scene_str,omitempty"`
				SceneID  int    `json:"scene_id,omitempty"`
			} `json:"scene"`
		}{
			Scene: struct {
				SceneStr string `json:"scene_str,omitempty"`
				SceneID  int    `json:"scene_id,omitempty"`
			}{
				SceneStr: req.GetSceneStr(),
				SceneID:  int(req.GetSceneId()),
			},
		},
	})
	if nil != err {
		s.logger.Errorf("WeChat GetQRTicket err, err:%v", err)
		return resp, err
	}

	resp.QrCode = basic.ShowQRCode(ticket)

	// 标记二维码扫描
	if err := s.wechatUsecase.SetQRCodeStatus(ctx, req.GetSceneStr(), time.Duration(req.GetExpireSeconds())*time.Second); nil != err {
		s.logger.Errorf("WeChat SetQRCodeStatus err, err:%v", err)
		return resp, err
	}

	return resp, nil
}

func (s *WeChatService) Check(rw ghttp.ResponseWriter, req *ghttp.Request) {
	m := make(map[string]string)
	mp := strings.FieldsFunc(req.URL.RawQuery, func(r rune) bool {
		if r == '=' || r == '&' {
			return true
		}
		return false
	})
	for i := 0; i < len(mp); i += 2 {
		m[mp[i]] = mp[i+1]
	}
	// rw.WriteHeader(200)
	rw.Write([]byte(m["echostr"]))
	return
}

func (s *WeChatService) WeChat(rw ghttp.ResponseWriter, req *ghttp.Request) error {

	// check
	// s.Check(rw, req)
	// return nil

	// 传入request和responseWriter
	server := s.officialAccount.GetServer(req, rw)
	// 设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		reply := new(message.Reply)
		switch msg.MsgType {
		case message.MsgTypeText:

		default:
			// TODO
			// 回复消息：演示回复用户发送的消息
			text := message.NewText(msg.Content)
			reply = &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}

		switch msg.Event {
		case message.EventScan, message.EventSubscribe:

			ctx := context.Background()
			// openId
			openId := msg.GetOpenID()

			// 清空二维码标识
			if err := s.wechatUsecase.ClearQRCodeStatus(ctx, msg.EventKey); nil != err {
				s.logger.Errorf("WeChat ClearQRCodeStatus err, err:%v", err)
				return nil
			}

			// 创建更新用户
			if err := s.CreateUpdateUser(ctx, openId); nil != err {
				s.logger.Errorf("WeChat CreateUpdateUser err, err:%v", err)
				return nil
			}

			return nil
		}

		return reply
	})

	// 处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		return err
	}
	// 发送回复的消息
	return server.Send()
}

// 创建更新用户
func (s *WeChatService) CreateUpdateUser(ctx context.Context, openId string) error {

	userCode := ""
	// 查询
	thirdPartUser, err := s.userUsecase.DetailByThird(ctx, openId, pbuser.ThirdPartType_LOGIN_TYPE_WECHAT)
	if nil != err {
		if !pberror.IsDataNotExist(err) {
			s.logger.Errorf("userUsecase DetailByThird err, err:%v", err)
			return err
		}
	} else {
		userCode = thirdPartUser.UserCode
	}

	if len(userCode) == 0 {
		// 添加
		newUserCode, err := s.userUsecase.Register(ctx, &biz.RegParam{
			ThirdPartUser: biz.ThirdPartUser{
				ThirdPartType: pbuser.ThirdPartType_LOGIN_TYPE_WECHAT,
				ThirdPartCode: openId,
			},
		})
		if nil != err {
			s.logger.Errorf("userUsecase Register err, err:%v", err)
			return err
		}
		userCode = newUserCode
	}

	// 微信用户信息
	userInfo, err := s.officialAccount.GetUser().GetUserInfo(openId)
	if err != nil {
		s.logger.Errorf("officialAccount GetUserInfo err, err:%v", err)
		return err
	}

	// 修改
	// todo  这儿应该把所用数据都记录一下
	if err := s.userUsecase.Update(ctx, &biz.User{
		UserCode:     userCode,
		WechatOpenid: openId,
		Portrait:     userInfo.Headimgurl,
		NickName:     userInfo.Nickname,
		Sex:          strconv.FormatInt(int64(userInfo.Sex), 10),
	}); nil != err {
		s.logger.Errorf("userUsecase Update err, err:%v", err)
		return err
	}

	return nil
}

func (s *WeChatService) CheckQRCode(ctx context.Context, req *pb.CheckQRCodeReq) (*pb.CheckQRCodeResp, error) {
	resp := new(pb.CheckQRCodeResp)

	ret, err := s.wechatUsecase.CheckQRCodeStatus(ctx, req.GetKey())
	if nil != err {
		return resp, err
	}

	resp.Result = ret

	return resp, nil
}
