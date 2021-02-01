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

func (a Article) Get(c *gin.Context) {}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
