package words2img

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/words2img/words2img_controllers"
)

func Router(r *gin.Engine) {
	words2img := r.Group("/words2img")
	{
		words2img.GET("/createImg", words2img_controllers.CreateImg)
	}
}
