package v1

import (
	"ginblog/constants/retCode"
	"ginblog/model"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv" // strconv 包实现了字符串与数字（整数、浮点数等）之间的互相转换
	"fmt"
)

// @description 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data) // 这一步可解析出 post 中的数据，存储到 data 变量中。
	
	code := model.AddArticle(&data)
	fmt.Println("API AddArticle data(res):", data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": retCode.GetRetCodeMsg(code),
		},
	)
}

// @description 查询单个文章
func QueryArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	var maps = make(map[string]interface{})
	articleInfo, code := model.QueryArticle(id)
	maps["title"] = articleInfo.Title
	maps["desc"] = articleInfo.Desc
	maps["content"] = articleInfo.Content
	maps["cover"] = articleInfo.Cover
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"message": retCode.GetRetCodeMsg(code),
		},
	)
}

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




