package service

import (
	"github.com/zxgit/gin-blog-project/internal/dao"
	"github.com/zxgit/gin-blog-project/internal/model"
)

type AddArticle struct {
	Title         string `json:"title"  binding:"required,max=100"`
	Summary       string `json:"summary" binding:"required"`
	CoverImageUrl string `json:"coverImageUrl" binding:"required"`
	Content       string `json:"content" binding:"required"`
	CreatedBy     string `json:"createdBy" binding:"required"`
	ModifiedBy    string `json:"modifiedBy" binding:"required"`
}

type GetArticleList struct {
	Page     int    `form:"page" binding:"min=1"`
	PageSize int    `form:"pageSize" binding:"required,min=10"`
	Category string `form:"category"`
}

func (s Service) AddArticle(params AddArticle) error {
	return s.Dao.AddArticle(&dao.Article{

		Title:         params.Title,
		Summary:       params.Summary,
		Content:       params.Content,
		CoverImageUrl: params.CoverImageUrl,
		CreatedBy:     params.CreatedBy,
		ModifiedBy:    params.ModifiedBy,
	})
}

func (s Service) GetArticleList(a GetArticleList) (list []model.Article, total int, err error) {
	page := a.Page
	pageSize := a.PageSize
	category := a.Category
	return s.Dao.GetArticleList(page, pageSize, category)
}
