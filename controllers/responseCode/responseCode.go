package responseCode

type ResCode int64

// A way to define a set of constants.
const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeRegister
	CodeServerBasy
)

// codeMsgMap A map, which is a key-value pair.
var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeRegister:        "注册失败",
	CodeServerBasy:      "服务器繁忙",
}

// Msg A method of ResCode.
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBasy]
	}
	return msg
}
