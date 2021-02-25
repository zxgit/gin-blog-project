package model

import (
	"github.com/zxgit/gin-blog-project/internal/dbClient"
	"github.com/zxgit/gin-blog-project/pkg/app"
)

type Focus struct {
	*Model
	Title         string `json:"title"`
	Img 		  string `json:"img"`
}

func (f Focus) TableName() string {
	return "blog_focus"
}

func (f Focus) GetFocusList(page, pageSize int) (list []Focus, err error) {
	//获取列表数
	err = dbClient.Db.Limit(pageSize).Offset(app.GetPageOffset(page, pageSize)).Find(&list).Error
	if err != nil {
		return list,  err
	}
	return
}
