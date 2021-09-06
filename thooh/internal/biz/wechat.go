package biz

import (
	"context"
	"time"
)

type WechatUsecase struct {
	repo WechatRepo
}

type WechatRepo interface {
	CheckQRCodeStatus(context.Context, string) (bool, error)
	SetQRCodeStatus(context.Context, string, time.Duration) error
	ClearQRCodeStatus(context.Context, string) error
}

func NewWechatUsecase(repo WechatRepo) *WechatUsecase {
	return &WechatUsecase{
		repo: repo,
	}
}

func (s *WechatUsecase) CheckQRCodeStatus(ctx context.Context, key string) (bool, error) {
	return s.repo.CheckQRCodeStatus(ctx, key)
}

func (s *WechatUsecase) SetQRCodeStatus(ctx context.Context, key string, expiration time.Duration) error {
	return s.repo.SetQRCodeStatus(ctx, key, expiration)
}

func (s *WechatUsecase) ClearQRCodeStatus(ctx context.Context, key string) error {
	return s.repo.ClearQRCodeStatus(ctx, key)
}
