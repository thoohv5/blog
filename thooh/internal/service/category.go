package service

import (
	"context"

	pbcategory "thooh/api/blog/v1/category"
	"thooh/internal/biz"
)

type CategoryService struct {
	pbcategory.UnimplementedCategoryServer
	usecase *biz.CategoryUsecase
}

func NewCategoryService(usecase *biz.CategoryUsecase) *CategoryService {
	return &CategoryService{
		usecase: usecase,
	}
}

func (s *CategoryService) Create(ctx context.Context, req *pbcategory.CreateReq) (*pbcategory.CreateResp, error) {
	resp := new(pbcategory.CreateResp)
	err := s.usecase.Create(ctx, &biz.Category{
		Label: req.GetLabel(),
		Name:  req.GetName(),
	})
	if nil != err {
		return resp, err
	}
	return resp, nil
}

func (s *CategoryService) List(ctx context.Context, req *pbcategory.ListReq) (*pbcategory.ListResp, error) {
	resp := new(pbcategory.ListResp)
	categoryList, err := s.usecase.List(ctx)
	if nil != err {
		return resp, err
	}

	for _, item := range categoryList {
		resp.List = append(resp.List, &pbcategory.Entity{
			Label: item.Label,
			Name:  item.Name,
		})
	}

	return resp, nil
}
