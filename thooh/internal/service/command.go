package service

import (
	"context"
	"strings"

	pb "thooh/api/common/v1/command"
	"thooh/internal/biz"
	"thooh/internal/conf"
)

type CommandService struct {
	pb.UnimplementedCommandServer
	file           *biz.File
	c              *conf.Assets
	articleUsecase *biz.ArticleUsecase
}

func NewCommandService(
	file *biz.File,
	c *conf.Assets,
	articleUsecase *biz.ArticleUsecase,
) *CommandService {
	return &CommandService{
		file:           file,
		c:              c,
		articleUsecase: articleUsecase,
	}
}

func (s *CommandService) Execute(ctx context.Context, req *pb.ExecuteReq) (resp *pb.ExecuteResp, err error) {
	resp = new(pb.ExecuteResp)

	switch req.GetCommand() {
	case "list":
		err = s.list(ctx)
	}

	return resp, err
}

func (s *CommandService) list(ctx context.Context) error {
	files, err := s.file.Walk(s.c.GetMd().GetDir())
	if nil != err {
		return err
	}
	for _, f := range files {
		if err := s.articleUsecase.Create(ctx, &biz.Article{
			Label: "10001",
			Name:  strings.TrimSuffix(f.Name, ".md"),
			Md:    f.Path,
		}); nil != err {
			return err
		}
	}
	return nil
}