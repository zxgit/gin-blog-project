package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/internal/service"
	"github.com/zxgit/gin-blog-project/pkg/app"
	"github.com/zxgit/gin-blog-project/pkg/errcode"
)

type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) GetToken(c *gin.Context) {
	response := app.NewResponse(c)
	srv := service.New(c)
	var tokenOpt service.GetTokenOpt
	valid, errs := app.BindAndValid(c, &tokenOpt)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	token, err := srv.GetToken(tokenOpt)
	if err != nil {
		response.ToErrorResponse(errcode.BusinessError.WithDetails(err.Error()))
		return
	}
	response.ToResponse(token)
	return
}
