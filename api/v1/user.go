package v1

import (
	"ginblog/constants/retCode"
	"ginblog/model"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv" // strconv 包实现了字符串与数字（整数、浮点数等）之间的互相转换
	"fmt"
)

// @description 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUserExist(data.Username)

	if code == retCode.SUCCESS {
		model.AddUser(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": retCode.GetRetCodeMsg(code),
		},
	)
}

// @description 删除用户
func DeleteUser(c *gin.Context) {
	// 字符串转int：Atoi()  int转字符串：Itoa()
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": retCode.GetRetCodeMsg(code),
		},
	)
}

// @description 更新用户信息
func UpdateUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Query("id"))
	_ = c.ShouldBindJSON(&data)
	fmt.Println("API UpdateUser data:", data)

	_, code := model.QueryUser(id)

	if code == retCode.SUCCESS {
		model.UpdateUser(id, &data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": retCode.GetRetCodeMsg(code),
		},
	)
}

// @description 查询单个用户
func QueryUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	var maps = make(map[string]interface{})
	userInfo, code := model.QueryUser(id)
	maps["username"] = userInfo.Username
	maps["role"] = userInfo.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"message": retCode.GetRetCodeMsg(code),
		},
	)
}

// @description 查询分页用户列表
func QueryUserList(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset")) // 偏移量
	size, _ := strconv.Atoi(c.Query("size")) // 单页大小
	fmt.Println("API QueryUser id:", offset, size)

	userList, hasMore, code := model.QueryUserList(offset, size)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    userList,
			"message": retCode.GetRetCodeMsg(code),
			"hasMore": hasMore,
		},
	)
}


