package model

type ArticleTag struct {
	*Model
	TagId     uint32 `json:"tagId"`
	ArticleId uint32 `json:"articleId"`
	CreatedBy string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}

func (t ArticleTag) TableName() string {
	return "blog_article_tag"
}
