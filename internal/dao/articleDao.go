package dao

import (
	"fmt"
	"github.com/zxgit/gin-blog-project/internal/dbClient"
	"github.com/zxgit/gin-blog-project/internal/model"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

/**添加
 */
func (d *Dao) AddArticle(params *Article) error{
	a := model.Article{
		Title: params.Title,
		Summary: params.Summary,
		CoverImageUrl: params.CoverImageUrl,
		Content: params.Content,
	}
	a.CreatedBy  = 1
	a.ModifiedBy = 1
	return  a.AddArticle()
}

/**
获取列表
 * @Author ZhangXin
 * @Description
 * @Date 15:28 2021/2/22
 * @Param
 * @return
 **/
func (d *Dao) GetArticleList(page,pageSize int,category string) (list []model.Article,total int,err error){
	a := model.Article{}
	var categoryModel model.Category
	categoryId := 0
	if category !="" {
		//todo 获取分类id
		dbClient.Db.Where("title=?",category).First(&categoryModel)
		categoryId = categoryModel.ID
		if categoryId <= 0 {
			return
		}
	}
	fmt.Println("分类id:",categoryId)
	return  a.GetArticleList(page,pageSize,categoryId)
}
