package service

import (
	"github.com/zxgit/gin-blog-project/internal/model"
)


/**
 * 获取焦点图数据
 * @Author ZhangXin
 * @Description
 * @Date 22:05 2021/2/22
 * @Param
 * @return
 **/
func (s Service) GetSiteInfo() (info model.SiteInfo,err error) {
	return s.Dao.GetSiteInfo()
}
