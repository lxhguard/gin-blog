# G006 gin路由

# 一.简介

这篇文章我们来使用一下 gin 框架的路由，同时去完成 User 的CURD接口。

这里可以下载一个软件 `API POST` ，我们去做接口的 CURD 请求。这是很好用的一个可视化工具。

[API POST官方网站](https://www.apipost.cn/)

# 二.gin router

[gin router group 路由组](https://gin-gonic.com/zh-cn/docs/examples/grouping-routes/)

根据上述教程，照猫画虎写代码。

```go
package router

import (
	"github.com/gin-gonic/gin"
	"ginblog/config"
)

func main() {
	gin.SetMode(config.AppMode)

	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/user/add", v1.AddUser)
	}

	router.Run(":8080")
}
```

这里我们使用了接口方法 `v1.AddUser`， 下一步我们就去开发接口。

# 三.接口

业务需求是需要迭代的，所以通常来说，我们使用 `v1`、`v2` 等来区分接口的版本迭代。

新建 `gin-blog/api/v1` 文件夹，去开发第一版本的接口逻辑。

```go
package v1

import (
	"ginblog/constants/retCode"
	"ginblog/model"

	"github.com/gin-gonic/gin"
	"net/http"

)

// AddUser 添加用户
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
```

接口就开发完毕了，然后使用 API-post 进行测试，自测OK。

接口的CURD，我们只完成了新增，接下来再去完成删除、修改、查询。

> 在新增的时候，其实已经完成了单条数据的查询。接下来去查询所有，以及查询分页。

删除接口如下：

```go
// model/User.go
// @description  删除用户 [GORM软删除](https://gorm.io/zh_CN/docs/delete.html)
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return retCode.ERROR_USER_DELETE
	}
	return retCode.SUCCESS
}

// api/v1/user.go
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

// router/router.go
func CreateRouter() {
	gin.SetMode(config.AppMode)

	r := gin.Default()

	// 简单的路由组: v1
	router := r.Group("/v1")
	{
		router.DELETE("/user/delete/:id", v1.DeleteUser)
	}

	r.Run(config.HttpPort)
}

// api post 软件测试路径:   
// DELETE请求
// 请求路径 localhost:3000/v1/user/delete/2
// 请求路径 localhost:3000/v1/user/delete/用户id
```



修改用户信息的接口如下：

```go
// model/User.go
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

// api/v1/user.go
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

// router/router.go
func CreateRouter() {
	gin.SetMode(config.AppMode)

	r := gin.Default()

	// 简单的路由组: v1
	router := r.Group("/v1")
	{
		router.POST("/user/update", v1.UpdateUser)
		// localhost:3000/v1/user/update?id=4 body&password=123456&username=csaj
	}

	r.Run(config.HttpPort)
}
```





查询单个用户信息的接口如下：

```go
// model/User.go
// @description 查询用户
func GetUser(id int) (User, int) {
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

// api/v1/user.go
// @description 查询单个用户
func QueryUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	var maps = make(map[string]interface{})
	userInfo, code := model.GetUser(id)
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
// router/router.go
func CreateRouter() {
	gin.SetMode(config.AppMode)

	r := gin.Default()

	// 简单的路由组: v1
	router := r.Group("/v1")
	{
		router.GET("/user/query", v1.QueryUser) // localhost:3000/v1/user/query?id=2
	}

	r.Run(config.HttpPort)
}
```


查询分页用户列表的接口如下：

```go
// model/User.go
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


// api/v1/user.go
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

// router/router.go
func CreateRouter() {
	gin.SetMode(config.AppMode)

	r := gin.Default()

	// 简单的路由组: v1
	router := r.Group("/v1")
	{
		router.GET("/user/queryall", v1.QueryUserList)
		// localhost:3000/v1/user/queryall?offset=0&size=10 偏移量 单页大小
	}

	r.Run(config.HttpPort)
}
```




