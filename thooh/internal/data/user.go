package data

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	pberror "thooh/api/error"
	"thooh/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

const (
	UserKey = "thooh:user"
)

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) Create(ctx context.Context, bu *biz.User) error {
	userHashKey := getUserHashKey(bu.UserCode)

	kv := make(map[string]interface{})
	bcVal := reflect.Indirect(reflect.ValueOf(bu))
	bcType := reflect.TypeOf(bu).Elem()
	for i := 0; i < bcVal.NumField(); i++ {
		kv[bcType.Field(i).Tag.Get("redis")] = bcVal.Field(i).Interface()
	}

	setHash := u.data.rdb.HMSet(ctx, userHashKey, kv)
	if err := setHash.Err(); nil != err {
		u.log.WithContext(ctx).Errorf("redis HMSet err, err:%v", err)
		return err
	}
	setList := u.data.rdb.LPush(ctx, getUserListKey(), userHashKey)
	if err := setList.Err(); nil != err {
		u.log.WithContext(ctx).Errorf("redis LPush err, err:%v", err)
		return err
	}
	return nil
}

func (u *userRepo) Update(ctx context.Context, bu *biz.User) error {
	userHashKey := getUserHashKey(bu.UserCode)
	// 旧值
	oldBu := new(biz.User)
	if err := u.data.rdb.HGetAll(ctx, userHashKey).Scan(oldBu); nil != err {
		if redis.Nil == err {
			return pberror.ErrorDataNotExist("redis HMGet err, err:%v", err)
		}
		u.log.WithContext(ctx).Errorf("redis HMGet err, err:%v", err)
		return err
	}

	// 比较
	kv := make(map[string]interface{})
	bcVal := reflect.Indirect(reflect.ValueOf(bu))
	oldBcVal := reflect.Indirect(reflect.ValueOf(oldBu))
	bcType := reflect.TypeOf(bu).Elem()
	for i := 0; i < bcVal.NumField(); i++ {
		if newBu := bcVal.Field(i).Interface(); !reflect.DeepEqual(newBu, oldBcVal.Field(i).Interface()) {
			kv[bcType.Field(i).Tag.Get("redis")] = newBu
		}
	}

	// 赋值
	setHash := u.data.rdb.HMSet(ctx, userHashKey, kv)
	if err := setHash.Err(); nil != err {
		u.log.WithContext(ctx).Errorf("redis HMSet err, err:%v", err)
		return err
	}

	return nil
}

func (u *userRepo) Detail(ctx context.Context, userCode string) (*biz.User, error) {
	user := new(biz.User)

	err := u.data.rdb.HGetAll(ctx, getUserHashKey(userCode)).Scan(user)
	if nil != err {
		if redis.Nil == err {
			return user, pberror.ErrorDataNotExist("redis HGet err, err:%v", err)
		}
		u.log.WithContext(ctx).Errorf("redis HGet err, err:%v", err)
		return user, err
	}

	return user, nil
}

func getUserHashKey(code string) string {
	return fmt.Sprintf("%v:%v", UserKey, code)
}

func getUserListKey() string {
	return fmt.Sprintf("%v:list", UserKey)
}
