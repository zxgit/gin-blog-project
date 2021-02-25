package service

import (
	"github.com/zxgit/gin-blog-project/internal/model"
)

type GetFocusList struct {
	Page         int `form:"page" binding:"min=1"`
	PageSize     int `form:"pageSize" binding:"required,max=5"`
}

/**
 * 获取焦点图数据
 * @Author ZhangXin
 * @Description
 * @Date 22:05 2021/2/22
 * @Param
 * @return
 **/
func (s Service) GetFocusList(a GetFocusList) (list []model.Focus,err error) {
	page := a.Page
	pageSize := a.PageSize
	return s.Dao.GetFocusList(page,pageSize)
}
