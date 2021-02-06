package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/internal/service"
	"github.com/zxgit/gin-blog-project/pkg/app"
	"github.com/zxgit/gin-blog-project/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) AddArticle() {

}


func (a Article) Get(c *gin.Context)  {}
func (a Article) List(c *gin.Context) {}

func (a Article) Create(c *gin.Context) {
	srv := service.New(c)
	var add service.AddArticle
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &add)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := srv.AddArticle(add)
	if err!=nil {
		response.ToErrorResponse(errcode.BusinessError.WithDetails(err.Error()))
	}
	response.ToResponse(gin.H{})
	return
}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
