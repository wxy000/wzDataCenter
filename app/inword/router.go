package inword

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/inword/inword_controllers"
)

func Router(r *gin.Engine) {
	inword := r.Group("/inword")
	{
		inword.GET("/getRandomWord", inword_controllers.GetRandomWord)
		inword.GET("getRandomImg", inword_controllers.GetRandomImg)
	}
}
