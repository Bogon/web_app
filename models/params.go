package models

// 定义请求参数结构体

// ParamSignUp `ParamSignUp` is a struct with three fields, `Username`, `Password`, and `RePassword`.
//
// The first line of the type definition is the type name, `ParamSignUp`. The type name is followed by a struct keyword,
// and then a list of fields inside curly braces `{}`. Each field has a name and a type, separated by a colon `:`.
//
// In Go, a name is exported if it begins with a capital letter. `ParamSignUp` is an exported type, `Username`, `Password`,
// and `RePassword
// @property {string} Username - The username of the user.
// @property {string} Password - The password of the user.
// @property {string} RePassword - The password that the user enters again to confirm the password.
type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
