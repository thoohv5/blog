package service

import (
	"context"

	pb "thooh/api/common/v1/user"
	"thooh/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	userCase *biz.UserUsecase
}

func NewUserService(userCase *biz.UserUsecase) *UserService {
	return &UserService{
		userCase: userCase,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	resp := new(pb.LoginResp)

	thirdPartUserCode := ""
	if userLoginParam := req.GetUserLoginParam(); nil != userLoginParam {
		thirdPartUserCode = userLoginParam.GetUserCode()
	}

	thirdPartUserType := pb.ThirdPartType_LOGIN_TYPE_DEFAULT
	if thirdPartLoginParam := req.GetThirdLoginParam(); nil != thirdPartLoginParam {
		thirdPartUserCode = thirdPartLoginParam.GetThirdPartCode()
		thirdPartUserType = thirdPartLoginParam.GetThirdPartType()
	}

	// 登陆
	token, err := s.userCase.Login(ctx, &biz.ThirdPartUser{
		ThirdPartType: thirdPartUserType,
		ThirdPartCode: thirdPartUserCode,
	})
	if nil != err {
		return resp, err
	}

	resp.Token = token

	return resp, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	resp := new(pb.RegisterResp)

	// 注册
	userCode, err := s.userCase.Register(ctx, &biz.RegParam{
		ThirdPartUser: biz.ThirdPartUser{
			ThirdPartType: req.GetThirdPartType(),
			ThirdPartCode: req.GetThirdPartCode(),
		},
		Password: req.GetPassword(),
	})
	if nil != err {
		return resp, err
	}

	resp.UserCode = userCode

	return resp, nil
}
func (s *UserService) Info(ctx context.Context, req *pb.InfoReq) (*pb.InfoResp, error) {
	resp := new(pb.InfoResp)

	// 用户信息
	userInfo, err := s.userCase.Detail(ctx, req.GetUserCode())
	if nil != err {
		return resp, err
	}

	resp = s.toUser(userInfo)

	return resp, nil
}

func (s *UserService) toUser(user *biz.User) *pb.InfoResp {
	return &pb.InfoResp{
		UserCode:     user.UserCode,
		Password:     user.Password,
		Phone:        user.Phone,
		Email:        user.Email,
		WechatOpenid: user.WechatOpenid,
		Portrait:     user.Portrait,
		NickName:     user.NickName,
		Sex:          user.Sex,
	}
}
