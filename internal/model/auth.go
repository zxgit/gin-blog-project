package model

import (
	"github.com/zxgit/gin-blog-project/internal/dbClient"
)

type Auth struct {
	*Model
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

type AuthInfo struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

/**
 * @Author ZhangXin
 * @Description //获取站点信息
 * @Date 11:14 2021/2/24
 * @return Site
 **/
func (a Auth) GetAuthInfo(appKey, appSecret string) (info AuthInfo, err error) {
	err = dbClient.Db.Where("app_key=? AND app_secret=?", appKey, appSecret).First(&a).Scan(&info).Error
	if err != nil {
		return info, err
	}
	return
}
