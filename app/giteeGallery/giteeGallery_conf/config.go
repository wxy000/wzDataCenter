package giteeGallery_conf

import (
	"gopkg.in/ini.v1"
	"log"
)

type Gitee struct {
	AccessToken string
	Branch      string
	Owner       string
	Repo        string
}

type GITEEGALLERY_CONF struct {
	Gitee Gitee
}

func GITEEGALLERY_InitConf(source string) *GITEEGALLERY_CONF {
	// 读取配置文件
	conf, err := ini.Load(source)
	if err != nil {
		log.Fatal("配置文件读取失败, err = ", err)
		return nil
	}
	cf := GITEEGALLERY_CONF{
		Gitee: Gitee{
			AccessToken: conf.Section("gitee").Key("access_token").String(),
			Branch:      conf.Section("gitee").Key("branch").String(),
			Owner:       conf.Section("gitee").Key("owner").String(),
			Repo:        conf.Section("gitee").Key("repo").String(),
		},
	}
	return &cf
}
