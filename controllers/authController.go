package controllers

import (
	"github.com/gin-gonic/gin"
	"wzDataCenter/common"
	"wzDataCenter/models"
	"wzDataCenter/utils"
)

// AuthHandler 登录
func AuthHandler(ctx *gin.Context) {
	// 用户发送用户名和密码过来
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// 校验用户名和密码是否正确
	user, err := models.LoginByUsername(username, password)
	if err != nil {
		common.FailWithMsg(err.Error(), ctx)
	} else {
		// 生成Token
		tokenString, _ := utils.NewJWT().GenToken(*user)
		common.OkWithData(gin.H{"token": tokenString}, ctx)
	}
}

// GetUserInfo 根据token获取用户信息
func GetUserInfo(ctx *gin.Context) {
	user, err := ctx.Get("claims")
	if !err {
		common.FailWithMsg("获取用户信息失败", ctx)
	} else {
		common.OkWithData(user, ctx)
	}
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(ctx *gin.Context) {
	user := models.Users{
		Username: ctx.PostForm("username"),
		Realname: ctx.PostForm("realname"),
		Phone:    ctx.PostForm("phone"),
		Email:    ctx.PostForm("email"),
		Gender:   ctx.PostForm("gender"),
		Mark:     ctx.PostForm("mark"),
	}
	if err := models.UpdateUserInfoByID(user); err != nil {
		common.FailWithMsg(err.Error(), ctx)
	} else {
		common.OkWithMsg("用户信息更新成功", ctx)
	}
}

// UpdatePassword 修改密码
func UpdatePassword(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userid := user.(models.Users).ID
	oldpw := ctx.PostForm("oldpassword")
	password := ctx.PostForm("password")
	if err := models.UpdatePasswordByID(userid, oldpw, password); err != nil {
		common.FailWithMsg(err.Error(), ctx)
	} else {
		common.OkWithMsg("密码更新成功", ctx)
	}
}
