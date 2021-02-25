package model

import (
	"github.com/zxgit/gin-blog-project/internal/dbClient"
	"github.com/zxgit/gin-blog-project/pkg/app"
)

type Category struct {
	*Model
	Title         string `json:"title"`
	Img 		  string `json:"img"`
}

func (c Category) TableName() string {
	return "blog_category"
}

func (c Category) GetCategoryList(page, pageSize int) (list []Category, err error) {
	//获取列表数
	err = dbClient.Db.Limit(pageSize).Offset(app.GetPageOffset(page, pageSize)).Find(&list).Error
	if err != nil {
		return list,  err
	}
	return
}
