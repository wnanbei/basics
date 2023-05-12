package code

// Code 请求状态码
type Code int

const (
	SUCCESS Code = 0 // SUCCESS 请求成功

	Failed            Code = -1 // Failed 请求失败
	ParamsParseFailed Code = -2 // ParamsParseFailed 参数解析失败

	UserNotFound  Code = -10000 // UserNotFound 用户不存在
	UserExisted   Code = -10001 // UserExisted 用户已存在
	PasswordError Code = -10002 // PasswordError 密码错误
)

var codeMsg = map[Code]string{
	SUCCESS: "ok",

	Failed:            "failed",
	ParamsParseFailed: "params parse failed",

	UserNotFound:  "user not found",
	PasswordError: "password error",
}

// GetMsg 获取请求状态码
func GetMsg(code Code) string {
	return codeMsg[code]
}

// Error 携带 code 的 error
type Error struct {
	Code Code
	Err  error
}

// NewError 构造 error
func NewError(code Code, err error) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

// Error ...
func (e Error) Error() string {
	return GetMsg(e.Code) + ": " + e.Err.Error()
}
