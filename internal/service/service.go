package service

import (
	"context"

	"blog-service/global"
	"blog-service/internal/dao"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func NewService(c context.Context) Service {
	return Service{
		ctx: c,
		dao: dao.NewDao(otgorm.WithContext(c, global.DBEngine)),
	}
}
