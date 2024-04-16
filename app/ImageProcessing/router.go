package ImageProcessing

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/ImageProcessing/ImageProcessing_controllers"
)

func Router(r *gin.Engine) {
	words2img := r.Group("/ImageProcessing")
	{
		words2img.GET("/createImg", ImageProcessing_controllers.CreateWordsImg)
	}
}
