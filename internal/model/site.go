package model

import (
	"github.com/zxgit/gin-blog-project/internal/dbClient"
)

type Site struct {
	*Model
	Avatar        string `json:"avatar"`
	Slogan 		  string `json:"slogan"`
	Name 		  string `json:"name"`
	Domain 		  string `json:"domain"`
	Notice 		  string `json:"notice"`
	Desc 		  string `json:"desc"`
}

type SiteInfo struct {
	Avatar        string `json:"avatar"`
	Slogan 		  string `json:"slogan"`
	Name 		  string `json:"name"`
	Domain 		  string `json:"domain"`
	Notice 		  string `json:"notice"`
	Desc 		  string `json:"desc"`
}

func (s Site) TableName() string {
	return "blog_site"
}

/**
 * @Author ZhangXin
 * @Description //获取站点信息
 * @Date 11:14 2021/2/24
 * @return Site
 **/
func (s Site) GetSiteInfo() (info SiteInfo, err error) {
	err = dbClient.Db.First(&s).Scan(&info).Error
	if err != nil {
		return info,  err
	}
	return
}
