package data

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	"thooh/internal/biz"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

const (
	ArticleKey = "thooh:article"
)

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (cr *articleRepo) Create(ctx context.Context, bc *biz.Article) error {
	articleHashKey := getArticleHashKey(bc.Code)

	kv := make(map[string]interface{})
	bcVal := reflect.Indirect(reflect.ValueOf(bc))
	bcType := reflect.TypeOf(bc).Elem()
	for i := 0; i < bcVal.NumField(); i++ {
		kv[bcType.Field(i).Tag.Get("redis")] = bcVal.Field(i).Interface()
	}

	setHash := cr.data.rdb.HMSet(ctx, articleHashKey, kv)
	if err := setHash.Err(); nil != err {
		cr.log.WithContext(ctx).Errorf("redis HMSet err, err:%v", err)
		return err
	}
	setList := cr.data.rdb.LPush(ctx, getArticleListKey(), articleHashKey)
	if err := setList.Err(); nil != err {
		cr.log.WithContext(ctx).Errorf("redis LPush err, err:%v", err)
		return err
	}
	return nil
}

func (cr *articleRepo) List(ctx context.Context, param *biz.ListParam) ([]*biz.Article, error) {
	resp := make([]*biz.Article, 0)
	getList := cr.data.rdb.LRange(ctx, getArticleListKey(), int64(param.Start), int64(param.Start)+int64(param.Limit)-1)

	ret, err := getList.Result()
	if nil != err {
		cr.log.WithContext(ctx).Errorf("redis LRange err,err:%v", err)
		return resp, err
	}

	if err := func() error {
		pipe := cr.data.rdb.Pipeline()
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
			article := new(biz.Article)
			if err := stringStringMapCmd.Scan(article); nil != err {
				return err
			}
			resp = append(resp, article)
		}
		return err
	}(); nil != err {
		cr.log.WithContext(ctx).Warnf("redis Pipeline err,err;%v", err)
		return resp, err
	}

	return resp, nil
}

func (cr *articleRepo) Detail(ctx context.Context, code string) (*biz.Article, error) {
	article := new(biz.Article)
	return article, cr.data.rdb.HGetAll(ctx, getArticleHashKey(code)).Scan(article)
}

func getArticleHashKey(code string) string {
	return fmt.Sprintf("%v:%v", ArticleKey, code)
}

func getArticleListKey() string {
	return fmt.Sprintf("%v:list", ArticleKey)
}
