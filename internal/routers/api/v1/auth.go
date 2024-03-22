package v1

import (
	"blog-service/global"
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.NewService(c.Request.Context())

	if err := svc.CheckAuth(&param); err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		errRsp := errcode.UnauthorizedAuthNotExist
		response.ToErrorResponse(errRsp)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		errRsp := errcode.UnauthorizedTokenGenerate
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
