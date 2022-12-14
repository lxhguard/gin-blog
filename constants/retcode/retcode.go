package retCode

const (
	SUCCESS = 200

	// User Error
	ERROR_USER_EXIST = 3001 // 用户已存在
	ERROR_USER_CREATE = 3002 // 用户创建失败
	ERROR_USER_DELETE = 3003 // 用户删除失败
	ERROR_USER_QUERY = 3004 // 用户查询失败
	ERROR_USER_NOT_EXIST = 3005 // 查询用户不存在
	ERROR_USER_QUERY_ALL = 3006 // 用户分页查询失败
	ERROR_USER_UPDATE = 3007 // 用户修改失败

	// Article Error
	ERROR_ARTICLE_CREATE = 3008 // 文章创建失败
	ERROR_ARTICLE_QUERY = 3009 // 文章查询失败
	ERROR_ARTICLE_EXIST = 3010 // 查询文章不存在
	ERROR_ARTICLE_QUERY_ALL = 3011 // 查询用户的所有文章失败
)

var retCodeMsg = map[int]string{
	SUCCESS: "SUCCESS",

	// User Error
	ERROR_USER_EXIST: "ERROR_USER_EXIST",
	ERROR_USER_CREATE: "ERROR_USER_CREATE",
	ERROR_USER_DELETE: "ERROR_USER_DELETE",
	ERROR_USER_QUERY: "ERROR_USER_QUERY",
	ERROR_USER_NOT_EXIST: "ERROR_USER_NOT_EXIST",
	ERROR_USER_QUERY_ALL: "ERROR_USER_QUERY_ALL",
	ERROR_USER_UPDATE: "ERROR_USER_UPDATE",

	// Article Error
	ERROR_ARTICLE_CREATE: "ERROR_ARTICLE_CREATE",
	ERROR_ARTICLE_QUERY: "ERROR_ARTICLE_QUERY",
	ERROR_ARTICLE_EXIST: "ERROR_ARTICLE_EXIST",
	ERROR_ARTICLE_QUERY_ALL: "ERROR_ARTICLE_QUERY_ALL",
}

func GetRetCodeMsg(code int) string {
	return retCodeMsg[code]
}