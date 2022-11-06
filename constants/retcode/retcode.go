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