package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/internal/service"
	"github.com/zxgit/gin-blog-project/pkg/app"
	"github.com/zxgit/gin-blog-project/pkg/errcode"
)

type Site struct {
}

func NewSite() Site {
	return Site{}
}

func (s Site) Info(c *gin.Context) {
	response := app.NewResponse(c)
	srv := service.New(c)
	info,err := srv.GetSiteInfo()
	if err!=nil {
		response.ToErrorResponse(errcode.BusinessError.WithDetails(err.Error()))
	}
	response.ToResponse(info)
	return
}

