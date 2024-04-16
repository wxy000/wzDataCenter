package ImageProcessing

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/ImageProcessing/ImageProcessing_controllers"
)

func Router(r *gin.Engine) {
	imageProcessing := r.Group("/imageProcessing")
	{
		imageProcessing.GET("/createWordsImg", ImageProcessing_controllers.CreateWordsImg)
		imageProcessing.POST("/createImgWaterMarkWithWordsAndIdio", ImageProcessing_controllers.CreateImgWaterMarkWithWordsAndIdio)
	}
}
