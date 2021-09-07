package data

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	"thooh/internal/biz"
)

type wechatRepo struct {
	data *Data
	log  *log.Helper
}

const (
	QRCheckKey = "thooh:qrCode:check"
	QRExtraKey = "thooh:qrCode:extra"
)

func NewWechatRepo(data *Data, logger log.Logger) biz.WechatRepo {
	return &wechatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (w *wechatRepo) CheckQRCodeStatus(ctx context.Context, key string) (bool, string, error) {
	exist, err := w.data.rdb.Exists(ctx, getQRCheckKey(key)).Result()
	if nil != err {
		return false, "", err
	}
	extra, err := w.data.rdb.Get(ctx, getQRExtrakey(key)).Result()
	if nil != err {
		if redis.Nil == err {
			return exist == 1, "", nil
		}
		return exist == 1, "", err
	}
	return exist == 1, extra, nil
}

func (w *wechatRepo) SetQRCodeStatus(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return w.data.rdb.Set(ctx, getQRCheckKey(key), value, expiration).Err()
}

func (w *wechatRepo) ClearQRCodeStatus(ctx context.Context, key string, extra string) error {
	w.data.rdb.Set(ctx, getQRExtrakey(key), extra, -1)
	return w.data.rdb.Del(ctx, getQRCheckKey(key)).Err()
}

func getQRCheckKey(key string) string {
	return fmt.Sprintf("%s:%s", QRCheckKey, key)
}

func getQRExtrakey(key string) string {
	return fmt.Sprintf("%s:%s", QRExtraKey, key)
}
