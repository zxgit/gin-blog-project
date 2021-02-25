package model

import (
	"github.com/zxgit/gin-blog-project/internal/dbClient"
	"github.com/zxgit/gin-blog-project/pkg/app"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	CoverImageUrl string `json:"coverImageUrl"`
	Content       string `json:"content"`
	state         uint8  `json:"state"`
	IsTop         int8   `json:"isTop"`
	IsHot         int8   `json:"isHot"`
	Summary       string `json:"summary"`
}

func (a Article) TableName() string {
	return "blog_article"
}

/**
 * @Author ZhangXin
 * @Description  添加文章
 * @Date 9:48 2021/2/24
 * @Param Article
 * @return
 **/
func (a Article) AddArticle() (err error) {
	err = dbClient.Db.Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

/**
 * @Author ZhangXin
 * @Description 获取文章列表
 * @Date 9:44 2021/2/24
 * @Param page 当前页数
 * @Param pageSize  每页大小
 * @Param categoryId 文章分类id
 * @return []Article, total int, err error
 **/
func (a Article) GetArticleList(page, pageSize int,categoryId int) (list []Article, total int, err error) {
	//获取列表数
	db := dbClient.Db.Limit(pageSize).Offset(app.GetPageOffset(page, pageSize))
	if categoryId > 0 {
		db = db.Where("category=? ",categoryId)
	}
	err = db.Find(&list).Error
	if err != nil {
		return list, 0, err
	}
	//获取总数
	dbCount:=dbClient.Db.Table(a.TableName())
	if categoryId > 0 {
		dbCount = dbCount.Where("category=? ",categoryId)
	}
	err = dbCount.Count(&total).Error
	if err != nil {
		return list, 0, err
	}
	return
}
