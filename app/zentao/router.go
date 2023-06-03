package zentao

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/zentao/zentao_common"
	"wzDataCenter/app/zentao/zentao_conf"
	"wzDataCenter/app/zentao/zentao_controllers"
	"wzDataCenter/middleware"
)

func Router(r *gin.Engine) {
	//**************初始化***************//
	// 读取配置文件
	zentao_common.ZENTAO_CONF = zentao_conf.ZENTAO_InitConf("./app/zentao/zentao_conf/config.ini")
	// 初始化数据库
	zentao_common.ZENTAO_DB = zentao_common.ZENTAO_InitDB(*zentao_common.ZENTAO_CONF)
	//**************初始化***************//

	zentao := r.Group("/zentao", middleware.JWTAuth())
	{
		zentao.GET("/getLeixingInfo", zentao_controllers.GetLeixingInfo)
	}
}
