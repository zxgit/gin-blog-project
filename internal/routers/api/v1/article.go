package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/internal/model"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

type use struct {
	tt int
	ww string
}
func (a Article) Get(c *gin.Context) {

}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {
	var articleAdd model.Article
	c.BindJSON(&articleAdd)
	articleAdd.AddArticle()
}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
