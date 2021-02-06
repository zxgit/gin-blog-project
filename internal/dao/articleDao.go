package dao

import (
	"github.com/zxgit/gin-blog-project/internal/model"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
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
		Desc: params.Desc,
		CoverImageUrl: params.CoverImageUrl,
		Content: params.Content,
		CreatedBy: params.Title,
		ModifiedBy: params.Title,
	}
	return  a.AddArticle()
}
