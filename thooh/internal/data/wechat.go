package data

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"thooh/internal/biz"
)

type wechatRepo struct {
	data *Data
	log  *log.Helper
}

const (
	QRCheckKey = "thooh:qrCode:check"
)

func NewWechatRepo(data *Data, logger log.Logger) biz.WechatRepo {
	return &wechatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (w *wechatRepo) CheckQRCodeStatus(ctx context.Context, key string) (bool, error) {
	exist, err := w.data.rdb.Exists(ctx, getQRCodeKey(key)).Result()
	return exist == 1, err
}

func (w *wechatRepo) SetQRCodeStatus(ctx context.Context, key string, expiration time.Duration) error {
	return w.data.rdb.Set(ctx, getQRCodeKey(key), true, expiration).Err()
}

func (w *wechatRepo) ClearQRCodeStatus(ctx context.Context, key string) error {
	return w.data.rdb.Del(ctx, getQRCodeKey(key)).Err()
}

func getQRCodeKey(key string) string {
	return fmt.Sprintf("%s:%s", QRCheckKey, key)
}
