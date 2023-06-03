package zentao_common

import (
	"gorm.io/gorm"
	"wzDataCenter/app/zentao/zentao_conf"
)

var ZENTAO_DB *gorm.DB
var ZENTAO_CONF *zentao_conf.ZENTAO_CONF
