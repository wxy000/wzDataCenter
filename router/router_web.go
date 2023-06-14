package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		web.GET("/hellohtml", func(ctx *gin.Context) {
			path := "请求路径：" + ctx.FullPath()
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"path":  path,
				"title": "Hello!",
			})
		})
	}
}
