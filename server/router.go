package server

import (
	"golang-crud/api"
	"golang-crud/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{





		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		guard := r.Group("/")
		guard.Use(middleware.AuthRequired())
		{
			// User Routing
			guard.GET("user/me", api.UserMe)
			guard.DELETE("user/logout", api.UserLogout)
		}


		//投
		v1.POST("videos",api.CreateVideo)

		// //find one
		v1.GET("videos/:id",api.ShowVideo)

		//find all
		v1.GET("videos",api.ListVideo)

		//更新
		v1.PUT("videos/:id",api.UpdateVideo)

		//删
		v1.DELETE("videos/:id",api.DeleteVideo)


	}
	return r
}
