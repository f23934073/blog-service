package service

import (
	"context"

	"blog-service/global"
	"blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func NewService(c context.Context) Service {
	return Service{
		ctx: c,
		dao: dao.NewDao(global.DBEngine),
	}
}
