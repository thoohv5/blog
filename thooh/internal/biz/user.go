package biz

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"

	pb "thooh/api/common/v1/user"
)

type User struct {
	// 用户标识
	UserCode string `redis:"userCode"`
	// 密码
	Password string `redis:"password"`

	// 电话
	Phone string `redis:"phone"`
	// 邮箱
	Email string `redis:"email"`
	// 微信
	WechatOpenid string `redis:"wechatOpenid"`

	// 头像
	Portrait string `redis:"portrait"`
	// 昵称
	NickName string `redis:"nickName"`
	// 性别
	Sex string `redis:"sex"`
}

type ThirdPartUser struct {
	// 用户标识
	UserCode string `redis:"userCode"`
	// 类型
	ThirdPartType pb.ThirdPartType `redis:"thirdPartType"`
	// 标识
	ThirdPartCode string `redis:"thirdPartCode"`
}

type UserRepo interface {
	Create(context.Context, *User) error
	Update(context.Context, *User) error
	Detail(context.Context, string) (*User, error)
}

type ThirdPartUserRepo interface {
	Create(context.Context, *ThirdPartUser) error
	Detail(context.Context, string, pb.ThirdPartType) (*ThirdPartUser, error)
}

type UserUsecase struct {
	logger            *log.Helper
	repo              UserRepo
	thirdPartUserRepo ThirdPartUserRepo
}

func NewUserUsecase(repo UserRepo, thirdPartUserRepo ThirdPartUserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, thirdPartUserRepo: thirdPartUserRepo, logger: log.NewHelper(logger)}
}

func (s *UserUsecase) Login(ctx context.Context, thirdPartUser *ThirdPartUser) (string, error) {
	userCode := ""
	if thirdPartUser.ThirdPartType == pb.ThirdPartType_LOGIN_TYPE_DEFAULT {
		userInfo, err := s.repo.Detail(ctx, thirdPartUser.ThirdPartCode)
		if nil != err {
			return userCode, err
		}
		userCode = userInfo.UserCode
	} else {
		thirdPartUser, err := s.thirdPartUserRepo.Detail(ctx, thirdPartUser.ThirdPartCode, thirdPartUser.ThirdPartType)
		if nil != err {
			return userCode, err
		}
		userCode = thirdPartUser.UserCode
	}

	return userCode, nil
}

type RegParam struct {
	// 第三方数据
	ThirdPartUser
	// 密码
	Password string
}

func (s *UserUsecase) Register(ctx context.Context, reg *RegParam) (string, error) {
	userCode := ""
	// 雪花算法
	node, err := snowflake.NewNode(1)
	if err != nil {
		return userCode, err
	}

	// Generate a snowflake ID.
	userCode = node.Generate().String()

	thirdPartCode := reg.ThirdPartCode
	thirdPartType := reg.ThirdPartType
	email := ""
	phone := ""
	wechatOpenid := ""
	switch thirdPartType {
	case pb.ThirdPartType_LOGIN_TYPE_EMAIL:
		email = thirdPartCode
	case pb.ThirdPartType_LOGIN_TYPE_PHONE:
		phone = thirdPartCode
	case pb.ThirdPartType_LOGIN_TYPE_WECHAT:
		wechatOpenid = thirdPartCode
	}

	// 用户
	if err := s.repo.Create(ctx, &User{
		UserCode:     userCode,
		Password:     reg.Password,
		Phone:        phone,
		Email:        email,
		WechatOpenid: wechatOpenid,
	}); nil != err {
		s.logger.Errorf("repo Create err, err:%v", err)
		return userCode, err
	}

	// 第三方
	if err := s.thirdPartUserRepo.Create(ctx, &ThirdPartUser{
		UserCode:      userCode,
		ThirdPartType: thirdPartType,
		ThirdPartCode: thirdPartCode,
	}); nil != err {
		s.logger.Errorf("thirdPartUserRepo Create err, err:%v", err)
		return userCode, err
	}

	return userCode, nil
}

func (s *UserUsecase) Update(ctx context.Context, user *User) error {
	// 用户
	if err := s.repo.Update(ctx, user); nil != err {
		s.logger.WithContext(ctx).Errorf("repo Create err, err:%v", err)
		return err
	}

	return nil
}

func (s *UserUsecase) Detail(ctx context.Context, userCode string) (*User, error) {
	return s.repo.Detail(ctx, userCode)
}

func (s *UserUsecase) DetailByThird(ctx context.Context, thirdPartCode string, thirdPartType pb.ThirdPartType) (*ThirdPartUser, error) {
	return s.thirdPartUserRepo.Detail(ctx, thirdPartCode, thirdPartType)
}
