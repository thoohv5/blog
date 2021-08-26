package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Category struct {
	// 唯一标识
	Label string `redis:"label"`
	// 名称
	Name string `redis:"name"`
}

type CategoryRepo interface {
	Create(context.Context, *Category) error
	List(context.Context) ([]*Category, error)
}

type CategoryUsecase struct {
	logger *log.Helper
	repo   CategoryRepo
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, logger: log.NewHelper(logger)}
}

func (c *CategoryUsecase) Create(ctx context.Context, bc *Category) error {
	return c.repo.Create(ctx, bc)
}

func (c *CategoryUsecase) List(ctx context.Context) ([]*Category, error) {
	return c.repo.List(ctx)
}
