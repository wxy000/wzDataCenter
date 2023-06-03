package zentao_conf

import (
	"gopkg.in/ini.v1"
	"log"
)

type Mysql struct {
	Alias   string
	Name    string
	User    string
	Pwd     string
	Host    string
	Port    string
	Charset string
}

type ZENTAO_CONF struct {
	Mysql Mysql
}

func ZENTAO_InitConf(source string) *ZENTAO_CONF {
	// 读取配置文件
	conf, err := ini.Load(source)
	if err != nil {
		log.Fatal("配置文件读取失败, err = ", err)
		return nil
	}
	cf := ZENTAO_CONF{
		Mysql: Mysql{
			Alias:   conf.Section("mysql").Key("db_alias").String(),
			Name:    conf.Section("mysql").Key("db_name").String(),
			User:    conf.Section("mysql").Key("db_user").String(),
			Pwd:     conf.Section("mysql").Key("db_pwd").String(),
			Host:    conf.Section("mysql").Key("db_host").String(),
			Port:    conf.Section("mysql").Key("db_port").String(),
			Charset: conf.Section("mysql").Key("db_charset").String(),
		},
	}
	return &cf
}
