package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/internal/service"
	"github.com/zxgit/gin-blog-project/pkg/app"
	"github.com/zxgit/gin-blog-project/pkg/errcode"
)

type Article struct {
}
type addArticleStruct struct {

}
func (a Article) AddArticle()  {

}

type addArticle struct {
	Title string `json:"title" binding:"required max:100"`
	Desc string  `json:"desc" binding:"required"`
	CoverImageUrl string `json:"desc" binding:"required"`
	Content string `json:"desc" binding:"required"`
	CreatedBy string `json:"desc" binding:"required"`
	ModifiedBy string `json:"desc" binding:"required"`
}
func (a Article) Get(c *gin.Context){}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {
	response := app.NewResponse(c)
	srv := service.New(c)
	var add addArticle
	valid,errs := app.BindAndValid(c,&add)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(nil)
	return
}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
