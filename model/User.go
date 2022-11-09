package model

import (
	"gorm.io/gorm"
	"ginblog/constants/retCode"
	"fmt"
)

type User struct {
	// [GORM模型定义 - 嵌入结构体](https://gorm.io/zh_CN/docs/models.html)
	gorm.Model
	// 账号(4 <= 长度 <= 8) 如果命名为 UserName, 则 mysql 列名为 user_name
	Username string `gorm:"type: varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"账号"`
	// 密码(6 <= 长度 <= 10) 
	Password string `gorm:"type: varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	// 权限等级 
	Role int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"权限等级"`
}

// @description  查询用户是否存在 [GORM查询](https://gorm.io/zh_CN/docs/query.html)
func CheckUserExist(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	fmt.Println("user:", user)
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

// @description 删除用户 [GORM软删除](https://gorm.io/zh_CN/docs/delete.html)
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return retCode.ERROR_USER_DELETE
	}
	return retCode.SUCCESS
}

// @description 更新用户信息 [GORM更新多列](https://gorm.io/zh_CN/docs/update.html)
func UpdateUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["password"] = data.Password
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return retCode.ERROR_USER_UPDATE
	}
	return retCode.SUCCESS
}

// @description 查询用户
func QueryUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, retCode.ERROR_USER_QUERY
	}

	if user.ID == 0 {
		return user, retCode.ERROR_USER_NOT_EXIST
	}

	return user, retCode.SUCCESS
}

// @description 查询用户列表
func QueryUserList(offset int, size int) ([]User, bool, int) {
	var userList []User
	var hasMore bool
	var lastQueryUser User

	err := db.Select("id, username, role, created_at").Offset(offset).Limit(size).Find(&userList).Error
	
	// TODO: hasMore 分页数据是否已经加载完毕
	// db.Select("id").Where("deleted_at = ''").Last(&lastQueryUser)
	// hasMore = lastQueryUser.ID == userList[len(userList)-1].ID
	fmt.Println("hasMore:", hasMore, lastQueryUser)

	if err != nil {
		return  userList, hasMore, retCode.ERROR_USER_QUERY_ALL
	}
	return  userList, hasMore, retCode.SUCCESS
}
