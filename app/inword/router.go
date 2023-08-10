package inword

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/inword/inword_controllers"
)

func Router(r *gin.Engine) {
	//设置html目录
	r.LoadHTMLGlob("./app/inword/web/html/*")

	inword := r.Group("/inword")
	{
		inword.GET("/getRandomWords", inword_controllers.GetRandomWords)
		inword.GET("/getRandomImgs", inword_controllers.GetRandomImgs)
	}
}
