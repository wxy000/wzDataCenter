package conf

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

type Jwt struct {
	Secret    string
	ExpiresAT string
	Issuer    string
}

type App struct {
	AppZentao string
}

type CONF struct {
	TokenName string
	HttpPort  string
	Mysql     Mysql
	Jwt       Jwt
	App       App
}

func InitConf(source string) *CONF {
	// 读取配置文件
	conf, err := ini.Load(source)
	if err != nil {
		log.Fatal("配置文件读取失败, err = ", err)
		return nil
	}
	cf := CONF{
		TokenName: conf.Section("").Key("token_name").String(),
		HttpPort:  conf.Section("").Key("http_port").String(),
		Mysql: Mysql{
			Alias:   conf.Section("mysql").Key("db_alias").String(),
			Name:    conf.Section("mysql").Key("db_name").String(),
			User:    conf.Section("mysql").Key("db_user").String(),
			Pwd:     conf.Section("mysql").Key("db_pwd").String(),
			Host:    conf.Section("mysql").Key("db_host").String(),
			Port:    conf.Section("mysql").Key("db_port").String(),
			Charset: conf.Section("mysql").Key("db_charset").String(),
		},
		Jwt: Jwt{
			Secret:    conf.Section("jwt").Key("secret").String(),
			ExpiresAT: conf.Section("jwt").Key("expiresat").String(),
			Issuer:    conf.Section("jwt").Key("issuer").String(),
		},
		App: App{
			AppZentao: conf.Section("app").Key("app_zentao").In("1", []string{"1", "0"}),
		},
	}
	return &cf
}
