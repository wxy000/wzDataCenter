package router

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/controllers"
	"wzDataCenter/middleware"
)

func Router(r *gin.Engine) {
	api := r.Group("/auth")
	{
		// 登录注册及设置相关
		api.POST("/login", controllers.AuthHandler)
		// 中间件
		api.Use(middleware.JWTAuth())
		// 获取用户信息
		api.GET("/getUserInfo", controllers.GetUserInfo)
		// 更新用戶信息
		api.POST("/updateUserInfo", controllers.UpdateUserInfo)
		// 更新密码
		api.POST("/updatePassword", controllers.UpdatePassword)

	}
}
