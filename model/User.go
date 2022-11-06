package model

import (
	"gorm.io/gorm"
	"ginblog/constants/retCode"
)

type User struct {
	// [GORM模型定义 - 嵌入结构体](https://gorm.io/zh_CN/docs/models.html)
	gorm.Model
	// 账号(4 <= 长度 <= 8)
	UserName string
	// 密码(6 <= 长度 <= 10)
	Password string
	// 权限等级
	Role int
}

// CheckUser 查询用户是否存在 [GORM查询](https://gorm.io/zh_CN/docs/query.html)
func CheckUserExist(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return retCode.ERROR_USER_EXIST
	}
	return retCode.SUCCESS
}

// @description  新增用户 [GORM创建](https://gorm.io/zh_CN/docs/create.html)
func AddUser(data *User) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return retCode.ERROR_USER_CREATE
	}
	return retCode.SUCCESS
}
