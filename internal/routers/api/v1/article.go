package v1

import (
	"github.com/gin-gonic/gin"
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
func (a Article) Get(c *gin.Context) {

}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
