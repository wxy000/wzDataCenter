package common

import (
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"wzDataCenter/conf"
)

const (
	ERROR      = -1
	ERRORLOGIN = 1001
	SUCCESS    = 0
)

var DB *gorm.DB
var CONF *conf.CONF
var CONFFILE *ini.File
