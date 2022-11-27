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
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// [GORM 迁移数据表](https://gorm.io/zh_CN/docs/migration.html)
	db.AutoMigrate(&User{}, &Article{})

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
