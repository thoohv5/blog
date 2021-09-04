package service

import (
	"context"
	ghttp "net/http"
	"strings"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"

	pb "thooh/api/wechat/v1/base"
	"thooh/internal/conf"
)

type WeChatService struct {
	pb.UnimplementedWeChatServer
	officialAccount *officialaccount.OfficialAccount
}

func NewWeChatService(c *conf.Wechat) *WeChatService {
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
		return resp, err
	}

	resp.QrCode = basic.ShowQRCode(ticket)
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
	// check(rw, req)
	// return

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
