package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/global"
	"strconv"
)

/**
 * @Author ZhangXin
 * @Description 获取页数大小
 * @Date 16:52 2021/1/24
 * @Param page
 * @return int
 **/
func GetPage(c *gin.Context) int {
	page ,_ := strconv.Atoi(c.Query("page"))
	if page < 0{
		return  1
	}
	return page
}

/**
 * @Author ZhangXin
 * @Description 获取每页数据量大小
 * @Date 16:52 2021/1/24
 * @Param  pageSize
 * @return int
 **/
func GetPageSize(c *gin.Context) int {
	pageSize ,_ := strconv.Atoi(c.Query("pageSize"))
	if pageSize <= 0{
		return  global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return  global.AppSetting.MaxPageSize
	}
	return pageSize
}

/**
 * @Author ZhangXin
 * @Description 获取偏移量
 * @Date 16:52 2021/1/24
 * @Param page     页数
 * @Param pageSize 页大小
 * @return int
 **/
func GetPageOffset(page,pageSize int) int{
	result := 0
	if page>0 {
		result = (page-1)*pageSize
	}
	return  result
}
