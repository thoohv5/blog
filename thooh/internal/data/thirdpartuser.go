package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	"thooh/api/common/v1/user"
	pberror "thooh/api/error"
	"thooh/internal/biz"
)

type thirdPartUserRepo struct {
	data *Data
	log  *log.Helper
}

const (
	ThirdPartUserKey = "thooh:user:third"
)

// NewThirdPartUserRepo .
func NewThirdPartUserRepo(data *Data, logger log.Logger) biz.ThirdPartUserRepo {
	return &thirdPartUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *thirdPartUserRepo) Create(ctx context.Context, btu *biz.ThirdPartUser) error {
	thirdPartUserHashKey := getThirdPartUserHashKey(btu.ThirdPartType)

	setHash := t.data.rdb.HSet(ctx, thirdPartUserHashKey, btu.ThirdPartCode, btu.UserCode)
	if err := setHash.Err(); nil != err {
		t.log.WithContext(ctx).Errorf("redis HMSet err, err:%v", err)
		return err
	}
	return nil
}

func (t *thirdPartUserRepo) Detail(ctx context.Context, thirdPartCode string, thirdPartType user.ThirdPartType) (*biz.ThirdPartUser, error) {
	thirdPartUser := new(biz.ThirdPartUser)
	thirdPartUser.ThirdPartCode = thirdPartCode
	thirdPartUser.ThirdPartType = thirdPartType

	err := t.data.rdb.HGet(ctx, getThirdPartUserHashKey(thirdPartType), thirdPartCode).Scan(&thirdPartUser.UserCode)
	if nil != err {
		if redis.Nil == err {
			return thirdPartUser, pberror.ErrorDataNotExist("redis HGet err, err:%v", err)
		}
		t.log.WithContext(ctx).Errorf("redis HGet err, err:%v", err)
		return thirdPartUser, err
	}

	return thirdPartUser, nil
}

func getThirdPartUserHashKey(thirdPartType user.ThirdPartType) string {
	return fmt.Sprintf("%v:%v", ThirdPartUserKey, thirdPartType)
}
