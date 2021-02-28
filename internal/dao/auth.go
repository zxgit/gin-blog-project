package dao

import (
	"errors"
	"github.com/zxgit/gin-blog-project/internal/model"
	"github.com/zxgit/gin-blog-project/pkg/app"
)

/**
 * 获取token
 * @Author ZhangXin
 * @Description
 * @Date 15:28 2021/2/22
 * @Param
 * @return model.Site
 **/
func (d *Dao) GetToken(appKey, appSecret string) (token string, err error) {
	//这里通过ak/sk的方式访问数据
	a := model.Auth{}
	info, err := a.GetAuthInfo(appKey, appSecret)
	if err != nil || info == (model.AuthInfo{}) {
		return "", errors.New("auth info does not exist.")
	}
	return app.GenerateToken(appKey, appSecret)
}
