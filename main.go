package main

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/app/inword"
	"wzDataCenter/app/parser12306"
	"wzDataCenter/app/zentao"
	"wzDataCenter/common"
	"wzDataCenter/conf"
	"wzDataCenter/middleware"
	"wzDataCenter/models"
	"wzDataCenter/router"
	"wzDataCenter/utils"
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
	// 黑名单
	r.Use(middleware.BlacklistMiddleware())
	// ip限流
	err := utils.SetupIPRateLimiter()
	if err != nil {
		return
	}
	r.Use(middleware.LimitMiddleware())
	//**************初始化***************//

	// 路由
	router.Router(r)
	/*子模块*/
	if common.CONF.App.AppZentao == "1" {
		zentao.Router(r)
	}
	if common.CONF.App.AppInword == "1" {
		inword.Router(r)
	}
	if common.CONF.App.Appparser12306 == "1" {
		parser12306.Router(r)
	}

	/*保存配置项到文件*/
	/*common.CONFFILE.Section("jwt").Key("issuer").SetValue("就你是懒猫啊")
	err := common.CONFFILE.SaveTo("./conf/config.ini")
	if err != nil {
		fmt.Println(err.Error())
	}*/

	// 运行
	port := ":" + common.CONF.HttpPort
	err1 := r.Run(port)
	if err1 != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
