package biz

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/russross/blackfriday/v2"
)

type Article struct {
	// 唯一标识
	Code string `redis:"code"`
	// 类目标识
	Label string `redis:"label"`
	// 标题
	Name string `redis:"name"`
	// 描述
	Desc string `redis:"desc"`
	// 内容
	Content string `redis:"content"`
	// 作者
	Author string `redis:"author"`
	// markdown
	Md string `redis:"md"`
	// 阅读数
	Read int32 `redis:"read"`
	// like
	Like int32 `redis:"like"`
}

type ArticleRepo interface {
	Create(context.Context, *Article) error
	List(context.Context, *ListParam) ([]*Article, error)
	Detail(context.Context, string) (*Article, error)
}

type ReadRepo interface {
	Increase(ctx context.Context, code string) error
}

type ArticleUsecase struct {
	logger   *log.Helper
	repo     ArticleRepo
	readRepo ReadRepo
}

func NewArticleUsecase(repo ArticleRepo, readRepo ReadRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, readRepo: readRepo, logger: log.NewHelper(logger)}
}

func (au *ArticleUsecase) Create(ctx context.Context, param *Article) error {

	// 没有内容
	if con := param.Content; con == "" {
		md := param.Md
		if md == "" {
			return errors.New("文章内容不存在")
		}

		if !checkFileIsExist(md) {
			return errors.New("文章地址不存在")
		}

		mdContent, err := au.parse(md)
		if nil != err {
			return err
		}
		param.Content = mdContent
	}

	if code := param.Code; code == "" {
		param.Code = fmt.Sprintf("%x", md5.Sum([]byte(param.Md)))
	}

	return au.repo.Create(ctx, param)
}

type ListParam struct {
	Start   int32
	Limit   int32
	KeyWord string
}

func (au *ArticleUsecase) List(ctx context.Context, param *ListParam) ([]*Article, error) {
	return au.repo.List(ctx, param)
}

func (au *ArticleUsecase) Detail(ctx context.Context, code string) (*Article, error) {
	resp := new(Article)
	if err := au.readRepo.Increase(ctx, code); nil != err {
		return resp, err
	}
	article, err := au.repo.Detail(ctx, code)
	if nil != err {
		return resp, err
	}
	resp = article
	return resp, nil
}

func (au *ArticleUsecase) parse(filename string) (string, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(blackfriday.Run(f)), nil
}

// 检查文件是否存在
func checkFileIsExist(filename string) bool {
	exist := false
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		exist = true
	}
	return exist
}