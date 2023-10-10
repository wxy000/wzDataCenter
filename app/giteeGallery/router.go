package giteeGallery

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/giteeGallery/giteeGallery_common"
	"wzDataCenter/app/giteeGallery/giteeGallery_conf"
	"wzDataCenter/app/giteeGallery/giteeGallery_controllers"
)

func Router(r *gin.Engine) {

	//**************初始化***************//
	// 读取配置文件
	giteeGallery_common.GITEEGALLERY_CONF = giteeGallery_conf.GITEEGALLERY_InitConf("./app/giteeGallery/giteeGallery_conf/config.ini")
	//**************初始化***************//

	giteeGallery := r.Group("/giteeGallery")
	{
		giteeGallery.POST("/update", giteeGallery_controllers.Update)
		giteeGallery.GET("/get", giteeGallery_controllers.Get)
		giteeGallery.GET("/del", giteeGallery_controllers.Del)
	}
}
