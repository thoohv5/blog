package data

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	"thooh/internal/biz"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

const (
	CategoryKey = "thooh:category"
)

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *categoryRepo) Create(ctx context.Context, bc *biz.Category) error {
	categoryHashKey := getCategoryHashKey(bc.Label)

	kv := make(map[string]interface{})
	bcVal := reflect.Indirect(reflect.ValueOf(bc))
	bcType := reflect.TypeOf(bc).Elem()
	for i := 0; i < bcVal.NumField(); i++ {
		kv[bcType.Field(i).Tag.Get("redis")] = bcVal.Field(i).Interface()
	}

	setHash := c.data.rdb.HMSet(ctx, categoryHashKey, kv)
	if err := setHash.Err(); nil != err {
		c.log.WithContext(ctx).Errorf("redis HMSet err, err:%v", err)
		return err
	}
	setList := c.data.rdb.LPush(ctx, getCategoryListKey(), categoryHashKey)
	if err := setList.Err(); nil != err {
		c.log.WithContext(ctx).Errorf("redis LPush err, err:%v", err)
		return err
	}
	return nil
}

func (c *categoryRepo) List(ctx context.Context) ([]*biz.Category, error) {
	resp := make([]*biz.Category, 0)
	getList := c.data.rdb.LRange(ctx, getCategoryListKey(), 0, -1)

	ret, err := getList.Result()
	if nil != err {
		c.log.WithContext(ctx).Errorf("redis LRange err,err:%v", err)
		return resp, err
	}

	if err := func() error {
		pipe := c.data.rdb.Pipeline()
		defer pipe.Close()
		cmdList := make([]*redis.StringStringMapCmd, 0)
		for _, item := range ret {
			cmdList = append(cmdList, pipe.HGetAll(ctx, item))
		}
		_, err := pipe.Exec(ctx)
		if nil != err {
			return err
		}
		for _, stringStringMapCmd := range cmdList {
			category := new(biz.Category)
			if err := stringStringMapCmd.Scan(category); nil != err {
				return err
			}
			resp = append(resp, category)
		}
		return err
	}(); nil != err {
		c.log.WithContext(ctx).Warnf("redis Pipeline err,err;%v", err)
		return resp, err
	}

	return resp, nil
}

func getCategoryHashKey(code string) string {
	return fmt.Sprintf("%v:%v", CategoryKey, code)
}

func getCategoryListKey() string {
	return fmt.Sprintf("%v:list", CategoryKey)
}
