package code

// Code 请求状态码
type Code int

const (
	SUCCESS Code = 0  // SUCCESS 请求成功
	Failed  Code = -1 // ERROR 请求失败
)

var codeMsg = map[Code]string{
	SUCCESS: "ok",
	Failed:  "failed",
}

// GetMsg 获取请求状态码
func GetMsg(code Code) string {
	return codeMsg[code]
}
