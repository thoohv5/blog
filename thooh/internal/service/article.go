package service

import (
	"context"

	pb "thooh/api/blog/v1/article"
	"thooh/internal/biz"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
	usecase *biz.ArticleUsecase
}

func NewArticleService(usecase *biz.ArticleUsecase) *ArticleService {
	return &ArticleService{
		usecase: usecase,
	}
}

func (s *ArticleService) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	return &pb.CreateResp{}, s.usecase.Create(ctx, &biz.Article{
		Code:    req.GetCode(),
		Label:   req.GetLabel(),
		Name:    req.GetName(),
		Desc:    req.GetDesc(),
		Content: req.GetContent(),
		Author:  req.GetAuthor(),
		Md:      req.GetMd(),
		Read:    req.GetRead(),
		Like:    req.GetLike(),
	})
}
func (s *ArticleService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	resp := new(pb.ListResp)

	resp.List = make([]*pb.ArticleEntity, 0)

	limit := req.GetLimit()

	articleList, err := s.usecase.List(ctx, &biz.ListParam{
		Start:   req.GetStart(),
		Limit:   limit + 1,
		KeyWord: req.GetKeyword(),
	})
	if nil != err {
		return resp, err
	}

	cur := len(articleList)
	if int32(cur) > limit {
		resp.HasMore = true
		resp.Start = req.GetStart() + req.GetLimit()
		articleList = articleList[0 : cur-1]
	}

	for _, item := range articleList {
		resp.List = append(resp.List, s.toEntity(item))
	}

	return resp, nil
}

func (s *ArticleService) Detail(ctx context.Context, req *pb.DetailReq) (*pb.DetailResp, error) {
	resp := new(pb.DetailResp)

	article, err := s.usecase.Detail(ctx, req.GetCode())
	if nil != err {
		return resp, err
	}

	resp.Entity = s.toEntity(article)

	return resp, nil
}

func (s *ArticleService) toEntity(item *biz.Article) *pb.ArticleEntity {
	return &pb.ArticleEntity{
		Code:    item.Code,
		Label:   item.Label,
		Name:    item.Name,
		Desc:    item.Desc,
		Content: item.Content,
		Author:  item.Author,
		Md:      item.Md,
		Read:    item.Read,
		Like:    item.Like,
	}
}
