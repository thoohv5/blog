package biz

import (
	"context"
	"time"
)

type WechatUsecase struct {
	repo WechatRepo
}

type WechatRepo interface {
	CheckQRCodeStatus(context.Context, string) (bool, string, error)
	SetQRCodeStatus(context.Context, string, interface{}, time.Duration) error
	ClearQRCodeStatus(context.Context, string, string) error
}

func NewWechatUsecase(repo WechatRepo) *WechatUsecase {
	return &WechatUsecase{
		repo: repo,
	}
}

func (s *WechatUsecase) CheckQRCodeStatus(ctx context.Context, key string) (bool, string, error) {
	return s.repo.CheckQRCodeStatus(ctx, key)
}

func (s *WechatUsecase) SetQRCodeStatus(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.repo.SetQRCodeStatus(ctx, key, value, expiration)
}

func (s *WechatUsecase) ClearQRCodeStatus(ctx context.Context, key string, extra string) error {
	return s.repo.ClearQRCodeStatus(ctx, key, extra)
}
