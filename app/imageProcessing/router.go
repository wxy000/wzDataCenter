package imageProcessing

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/imageProcessing/imageProcessing_controllers"
)

func Router(r *gin.Engine) {
	imageProcessing := r.Group("/imageProcessing")
	{
		imageProcessing.GET("/createWordsImg", imageProcessing_controllers.CreateWordsImg)
		imageProcessing.POST("/createImgWaterMarkWithWords", imageProcessing_controllers.CreateImgWaterMarkWithWords)
		imageProcessing.POST("/createImgWaterMarkWithIdio", imageProcessing_controllers.CreateImgWaterMarkWithIdio)
		imageProcessing.POST("/createImgWaterMarkJSON", imageProcessing_controllers.CreateImgWaterMarkJSON)
		imageProcessing.POST("/createImgWaterMarkFORM", imageProcessing_controllers.CreateImgWaterMarkFORM)
	}
}
