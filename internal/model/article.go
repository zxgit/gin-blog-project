package model

import "fmt"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"coverImageUrl"`
	Content       string `json:"content"`
	state         uint8  `json:"state"`
	CreatedBy     string `json:"createdBy"`
	ModifiedBy    string `json:"modifiedBy"`
}

func (t Article) TableName() string {
	return "blog_article"
}

func( t Article) AddArticle() bool{
	fmt.Println(t)
	db.Create(&t)
	return  true
}

