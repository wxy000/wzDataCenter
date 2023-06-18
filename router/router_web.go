package router

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/controllers/controllers_web"
)

func RouterWeb(r *gin.Engine) {
	//设置html目录
	r.LoadHTMLGlob("./web/html/*")
	//静态文件目录
	r.Static("/img", "./web/img")
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")
	r.Static("/layui", "./web/layui")

	web := r.Group("/web")
	{
		web.GET("/index", controllers_web.LayoutMenu)
	}
}
