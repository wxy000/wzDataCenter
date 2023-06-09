package main

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/zentao"
	"wzDataCenter/common"
	"wzDataCenter/conf"
	"wzDataCenter/middleware"
	"wzDataCenter/models"
)

func main() {

	// 加载banner
	common.ReadBanner("./common/banner.txt")

	r := gin.Default()

	//**************初始化***************//
	// 读取配置文件
	common.CONF = conf.InitConf("./conf/config.ini")
	// 初始化数据库
	common.DB = common.InitDB(*common.CONF)
	// 数据库迁移
	models.Setup()
	// 放行所有跨域请求
	r.Use(middleware.Cors())
	//**************初始化***************//

	// 路由
	router(r)
	/*子模块*/
	if common.CONF.App.AppZentao == "1" {
		zentao.Router(r)
	}

	// 运行
	port := ":" + common.CONF.HttpPort
	err1 := r.Run(port)
	if err1 != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
