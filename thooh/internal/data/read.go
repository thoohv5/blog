package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"thooh/internal/biz"
)

type ReadRepo struct {
	data *Data
	log  *log.Helper
}

// NewReadRepo .
func NewReadRepo(data *Data, logger log.Logger) biz.ReadRepo {
	return &ReadRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (rr *ReadRepo) Increase(ctx context.Context, code string) error {

	read, err := rr.data.rdb.HGet(ctx, getArticleHashKey(code), "read").Int64()
	if nil != err {
		return err
	}
	read++
	err = rr.data.rdb.HSet(ctx, getArticleHashKey(code), "read", read).Err()
	if nil != err {
		return err
	}
	return nil
}
