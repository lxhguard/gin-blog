# G007 gorm关联表

# 一.简介

这一小文章呢，我们去学习「关联表」。

之前的文章都是创建的 User 模型，每个用户都会进行发文章，所以这次我们创建一个 Article 模型。

Article 模型会关联  User 模型 ，最终实现一个「查询一个用户发表的所有文章」接口。

这里我真实使用的方式是：分别查询两个表，然后进行数据的组合。

实现联合的话，我这里总是有BUG，我排查不出来自己哪里写错了。再加上问了我兄弟，小马说之前在腾讯、字节、虾皮也几乎不用联合，基本都是组装数据，因为两者效率其实差不了太多。

# 二.Article 模型

```go
// 文章[关联 用户ID]
package model

import (
	"gorm.io/gorm"
	"ginblog/constants/retCode"
	"fmt"
)

type Article struct {
	// 关联 User 模型 [Belongs To](https://gorm.io/zh_CN/docs/belongs_to.html)
	User User `gorm:"foreignkey:Uid"` // GORM 同时提供自定义外键名字的方式
	// 外键，用于在 User 和 Article 之间创建一个外键关系
	Uid int `gorm:json:"uid" label:"当前文章所属于的用户ID"`

	// [GORM模型定义 - 嵌入结构体](https://gorm.io/zh_CN/docs/models.html)
	gorm.Model
	// 文章标题
	Title string `gorm:"type: varchar(20);not null" json:"title" label:"文章标题"`
	// 文章描述
	Desc string `gorm:"type: varchar(100);not null" json:"desc" label:"文章描述"`
	// 文章内容
	Content string `gorm:"type: longtext" json:"content" label:"文章内容"`
	// 文章封面图片
	Cover string `gorm:"type: varchar(300)" json:"cover" label:"文章封面图片"`
}
```

剩下的 api 部分 和 router 部分，跟着 User 照猫画虎就行。这里就不列举了。

建议好好学习 User 模型，这是用 go 开发业务的基础demo。

# 三.组装数据

```go
package v1
// @description 查询某个用户发表的所有文章列表
func QueryUserArticleAllList(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))

	articleList, code := model.QueryUserArticleAllList(uid)

	user, queryCode := model.QueryUser(uid)
	for i,_ := range articleList {
		articleList[i].User = user
	}
	if (queryCode != retCode.SUCCESS) {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  queryCode,
				"message": retCode.GetRetCodeMsg(queryCode),
			},
		)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    articleList,
			"message": retCode.GetRetCodeMsg(code),
		},
	)
}



package model
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

```









