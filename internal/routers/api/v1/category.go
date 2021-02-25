package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/internal/service"
	"github.com/zxgit/gin-blog-project/pkg/app"
	"github.com/zxgit/gin-blog-project/pkg/errcode"
)

type Category struct {
}

func NewCategory() Category {
	return Category{}
}

func (ca Category) List(c *gin.Context) {
	response := app.NewResponse(c)
	srv := service.New(c)
	var listOpt service.GetGetCategoryList
	valid, errs := app.BindAndValid(c, &listOpt)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	list,err := srv.GetGetCategoryList(listOpt)
	if err!=nil {
		response.ToErrorResponse(errcode.BusinessError.WithDetails(err.Error()))
	}
	response.ToResponseList(list,0)
	return
}

