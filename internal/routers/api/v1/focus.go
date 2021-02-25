package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/internal/service"
	"github.com/zxgit/gin-blog-project/pkg/app"
	"github.com/zxgit/gin-blog-project/pkg/errcode"
)

type Focus struct {
}

func NewFocus() Focus {
	return Focus{}
}

func (f Focus) List(c *gin.Context) {
	response := app.NewResponse(c)
	srv := service.New(c)
	var listOpt service.GetFocusList
	valid, errs := app.BindAndValid(c, &listOpt)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	list,err := srv.GetFocusList(listOpt)
	if err!=nil {
		response.ToErrorResponse(errcode.BusinessError.WithDetails(err.Error()))
	}
	response.ToResponseList(list,0)
	return
}

