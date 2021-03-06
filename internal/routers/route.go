package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zxgit/gin-blog-project/global"
	"github.com/zxgit/gin-blog-project/internal/middleware"
	v1 "github.com/zxgit/gin-blog-project/internal/routers/api/v1"
	"github.com/zxgit/gin-blog-project/pkg/limiter"
	"net/http"
	"time"
)

//对鉴权方法设置令牌桶
var methodLimiters = limiter.NewMethodLimiter().AddBucket(
	limiter.LimiterBucketRule{
		Key:          "/token",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.Recovery())
	//注册令牌桶
	r.Use(middleware.RateLimiter(methodLimiters))

	//统一超时60秒
	r.Use(middleware.ContextTimeout(global.AppSetting.Timeout))

	//添加允许跨域
	r.Use(Cors())
	//添加验证国际化
	r.Use(middleware.Translations())
	r.Use(middleware.AccessLog())
	article := v1.NewArticle()
	focus := v1.NewFocus()
	category := v1.NewCategory()
	site := v1.NewSite()
	auth := v1.NewAuth()
	//jwd授权
	r.POST("/token", auth.GetToken) //获取token
	apiV1 := r.Group("/api/v1")
	{
		//添加jwt token校验
		apiV1.Use(middleware.Jwt())
		//标签相关路由
		apiV1.GET("/focus", focus.List)        //焦点图列表
		apiV1.GET("/categorys", category.List) //文章分类列表
		apiV1.GET("/site", site.Info)          //站点信息

		//文章相关路由
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
