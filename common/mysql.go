package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"wzDataCenter/conf"
)

func InitDB(conf conf.CONF) *gorm.DB {
	var db *gorm.DB
	var err error

	// Mysql配置信息
	mysqlName := conf.Mysql.Name
	mysqlUser := conf.Mysql.User
	mysqlPwd := conf.Mysql.Pwd
	mysqlHost := conf.Mysql.Host
	mysqlPort := conf.Mysql.Port
	mysqlCharset := conf.Mysql.Charset

	var dataSource string
	dataSource = mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlName +
		"?charset=" + mysqlCharset + "&parseTime=True&loc=Local&multiStatements=true"
	db, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		err := db.Error
		if err != nil {
			return nil
		}
	}

	sqlDB, _ := db.DB()

	// 设置连接池，空闲连接
	sqlDB.SetMaxIdleConns(50)
	// 打开链接
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
