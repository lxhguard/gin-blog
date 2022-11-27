// 文章(关联 用户ID) [golang-gorm实现关联查询(四) ](https://www.cnblogs.com/davis12/p/16365294.html)
package model

import (
	"gorm.io/gorm"
	"ginblog/constants/retCode"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	// 关联 User 模型 [Belongs To](https://gorm.io/zh_CN/docs/belongs_to.html)
	// User User `gorm:"primary_key:Uid;"`
	User User `gorm:"foreignkey:Uid"` // GORM 同时提供自定义外键名字的方式
	// 外键，用于在 User 和 Article 之间创建一个外键关系
	Uid int `gorm:json:"uid" label:"当前文章所属于的用户ID"`

	// [GORM模型定义 - 嵌入结构体](https://gorm.io/zh_CN/docs/models.html)
	gorm.Model
	// 文章标题
	Title string `gorm:"type: varchar(20);not null" json:"title" label:"文章标题"`
	// 文章描述
	Desc string `gorm:"type: varchar(100);" json:"desc" label:"文章描述"`
	// 文章内容
	Content string `gorm:"type: longtext" json:"content" label:"文章内容"`
	// 文章封面图片
	Cover string `gorm:"type: varchar(300)" json:"cover" label:"文章封面图片"`
}


// @description  新增文章 [GORM创建](https://gorm.io/zh_CN/docs/create.html)
func AddArticle(data *Article) (code int) {
	// 通过Profile关联查询user数据, 查询结果保存到user变量
	// db.Model(&data).Association("User").Find(&data)

	err := db.Create(&data).Error
	fmt.Println("MODEL AddArticle err:", err)

	if err != nil {
		return retCode.ERROR_ARTICLE_CREATE
	}
	return retCode.SUCCESS
}

// @description 查询单个文章的数据
func QueryArticle(id int) (Article, int) {
	var article Article
	err := db.Limit(1).Where("id = ?", id).Find(&article).Error
	if err != nil {
		return article, retCode.ERROR_ARTICLE_QUERY
	}

	if article.ID == 0 {
		return article, retCode.ERROR_ARTICLE_EXIST
	}

	return article, retCode.SUCCESS
}

// @description 查询某个用户发表的所有文章列表
func QueryUserArticleAllList(uid int) ([]Article, int) {
	var articleList []Article
	err := db.Select("uid", "title", "content", "desc").Where("uid = ?", uid).Find(&articleList).Error
		// GORM设计问题：如果列名是 desc ，必须 `desc` ，或者单拎出去写 "desc"，或者使用 Order
	// err := db.Select("uid", "title", "content", "desc").Where("uid = ?", uid).Find(&articleList).Error
	// err := db.Select("uid, title, content, `desc`").Where("uid = ?", uid).Find(&articleList).Error
	// err := db.Select("uid, title, content").Where("uid = ?", uid).Order("id desc").Find(&articleList).Error

	if err != nil {
		return  articleList, retCode.ERROR_ARTICLE_QUERY_ALL
	}

	return  articleList, retCode.SUCCESS
}



