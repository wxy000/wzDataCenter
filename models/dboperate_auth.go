package models

import (
	"errors"
	"wzDataCenter/common"
)

// GetUserInfoByID 根据id获取用户信息
func GetUserInfoByID(userID uint) (int64, *Users, error) {
	var user Users
	if userID == 0 {
		return 0, nil, errors.New("用户编号不得为空")
	}
	result := common.DB.Where("userid = ?", userID).Limit(1).Find(&user)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	if result.RowsAffected < 1 {
		return 0, nil, errors.New("用户不存在")
	}
	return result.RowsAffected, &user, nil
}

// GetUserInfoByUsername 根据username获取用户信息
func GetUserInfoByUsername(username string) (int64, *Users, error) {
	var user Users
	if username == "" {
		return 0, nil, errors.New("用户编号不得为空")
	}
	result := common.DB.Where("username = ?", username).Limit(1).Find(&user)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	if result.RowsAffected < 1 {
		return 0, nil, errors.New("用户不存在")
	}
	return result.RowsAffected, &user, nil
}

// LoginByUsername 用户登录
func LoginByUsername(username string, password string) (*Users, error) {
	var user Users
	result := common.DB.Where("username = ? and password = ?", username, password).Limit(1).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New("用户名密码错误")
	}
	return &user, nil
}

// UpdateUserInfoByID 根据id更新用户信息
func UpdateUserInfoByID(user Users) error {
	result := common.DB.Model(&user).Omit("userid", "password").Where("userid = ?", user.ID).Updates(user)
	if result.RowsAffected < 1 {
		return result.Error
	}
	return nil
}

// UpdatePasswordByID 根据id修改用户密码
func UpdatePasswordByID(userid uint, oldpw string, password string) error {
	var user Users
	r1 := common.DB.Where("userid = ? and password = ?", userid, oldpw).First(&user)
	if r1.RowsAffected < 1 {
		return errors.New("旧密码错误")
	}
	if oldpw == password {
		return errors.New("新密码不得与旧密码相同")
	}
	r2 := common.DB.Model(&user).Where("userid = ?", userid).Update("password", password)
	if r2.RowsAffected < 1 {
		return errors.New("密码修改失败，请稍后再试")
	}
	return nil
}
