# G005 gorm使用

# 一.安装

这次我们来链接 mysql 数据库，可视化工具使用 workbench 就行。

gorm 是 Golang 目前比较热门的数据库ORM操作库，把struct类型和数据库表记录进行映射，操作数据库的时候不需要直接手写Sql代码。

先安装 gorm，如下：

```js
// mysql 驱动包
$ go get -u gorm.io/driver/mysql
// 安装 gorm 包
$ go get -u gorm.io/gorm
```

# 二.连接数据库

[GORM 官方文档](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)

我们先来建立一下 服务 和 mysql数据库 的连接。

```go
// gin-blog/model/db.go
package model

import (
	"ginblog/config"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	"gorm.io/gorm/schema"
    "fmt"
	"time"
)

var db *gorm.DB
var err error

// @description 连接 mysql 数据库
func ConnectDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DbUser, config.DbPassWord, config.DbHost, config.DbPort, config.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// [GORM 迁移数据表](https://gorm.io/zh_CN/docs/migration.html)
	db.AutoMigrate(&User{})

	// [GORM 连接池教程](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Printf("数据库连接成功: %s ", dsn)
}

```

主文件，如下：

```go
// gin-blog/main.go
package main

import (
	"ginblog/model"
)

func main() {
	model.ConnectDb()
}
```

# 三.建表

数据库链接好了之后，需要进行建表。

这次先创建 User 表，即 用户信息表。相当于一个 Schema 。

同时收拢了一些该表下的 CURD 原子方法，为 gin 路由接口服务。

```go
package model

import (
	"gorm.io/gorm"
	"ginblog/constants/retCode"
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

```

# 四.响应状态码

创建响应码，接口返回用。下一篇文章我们来写接口。

```go
package retCode

const (
	SUCCESS = 200

	// User Error
	ERROR_USER_EXIST = 3001 // 用户已存在
	ERROR_USER_CREATE = 3002 // 用户创建失败
)

var retCodeMsg = map[int]string{
	SUCCESS: "SUCCESS",
	ERROR_USER_CREATE: "ERROR_USER_CREATE",
}

func GetRetCodeMsg(code int) string {
	return retCodeMsg[code]
}
```














