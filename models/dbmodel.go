package models

import (
	"gorm.io/gorm"
	"wzDataCenter/common"
)

// Users 用户表
type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"size:20;not null;unique;comment:用户名"`
	Realname string `json:"realname" gorm:"size:100;not null;comment:真实姓名"`
	Password string `json:"-" gorm:"size:100;not null;default:123456;comment:密码"`
	Phone    string `json:"phone" gorm:"size:100;comment:手机号码"`
	Email    string `json:"email" gorm:"size:100;comment:邮箱"`
	Gender   string `json:"gender" gorm:"size:1;not null;default:m;comment:性别-m男w女"`
	Role     string `json:"role" gorm:"size:1;not null;defalut:1;comment:角色-0管理员1普通用户..."`
	Avat     string `json:"avat" gorm:"size:255;comment:头像"`
	Status   string `json:"status" gorm:"size:1;not null;default:1;comment:状态-0未激活1已激活"`
	Mark     string `json:"mark" gorm:"size:1000;comment:备注"`
}

func Setup() {
	autoMigrate(&Users{})
}

// 自动迁移
func autoMigrate(tables ...interface{}) {
	// 创建表时添加后缀
	common.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(tables...)
	// AutoMigrate 会创建表、缺失的外键、约束、列和索引。
	// 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。
}
