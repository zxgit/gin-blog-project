package service

import (
	"github.com/zxgit/gin-blog-project/internal/dao"
)


type AddArticle struct {
	Title         string `json:"title" binding:"required,max=100"`
	Desc          string `json:"desc" binding:"required"`
	CoverImageUrl string `json:"coverImageUrl" binding:"required"`
	Content       string `json:"content" binding:"required"`
	CreatedBy     string `json:"createdBy" binding:"required"`
	ModifiedBy    string `json:"modifiedBy" binding:"required"`
}

func (s Service) AddArticle(params AddArticle) error {
	return s.Dao.AddArticle(&dao.Article{
		Title:         params.Title,
		Desc:          params.Desc,
		Content:       params.Content,
		CoverImageUrl: params.CoverImageUrl,
		CreatedBy:     params.CreatedBy,
		ModifiedBy:    params.ModifiedBy,
	})
}
