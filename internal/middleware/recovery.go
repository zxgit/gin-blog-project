/*
@Time : 2021/3/1 22:38
@Author : zhangxin
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/global"
	"github.com/zxgit/gin-blog-project/pkg/app"
	"github.com/zxgit/gin-blog-project/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err:%v"
				global.Logger.WithCallersFrames().Errorf(s, err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.IsAborted()
			}
		}()
		c.Next()
	}
}
