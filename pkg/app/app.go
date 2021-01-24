/*
@Time : 2021/1/24 16:56
@Author : zhangxin
*/
package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}
type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
	TotalRows int `json:"totalRows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

/**
 * @Author ZhangXin
 * @Description 返回详情
 * @Date 17:00 2021/1/24
 * @Param
 * @return
 **/
func (r *Response)ToResponse(data interface{})  {
	if data==nil{
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK,data)
}

/**
 * @Author ZhangXin
 * @Description 返回列表数据
 * @Date 17:04 2021/1/24
 * @Param
 * @return
 **/
func (r *Response)ToResponseList(list interface{},totalRows int)  {
	r.Ctx.JSON(http.StatusOK,gin.H{
		"list":list,
		"pager": Pager{
			Page:     GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response)ToErrorResponse(err *errcode.Error)  {
	response :=gin.H{"code":err.Code(),"msg":err.Msg()}
	details :=err.Details()
	if len(details)>0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(),response)
}

