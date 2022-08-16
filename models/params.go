package models

// 定义请求参数结构体

// ParamSignUp `ParamSignUp` is a struct with three fields, `Username`, `Password`, and `RePassword`.
//	注册填写信息
// The first line of the type definition is the type name, `ParamSignUp`. The type name is followed by a struct keyword,
// and then a list of fields inside curly braces `{}`. Each field has a name and a type, separated by a colon `:`.
//
// In Go, a name is exported if it begins with a capital letter. `ParamSignUp` is an exported type, `Username`, `Password`,
// and `RePassword
// @property {string} Username - The username of the user.
// @property {string} Password - The password of the user.
// @property {string} RePassword - The password that the user enters again to confirm the password.
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin `ParamLogin` is a struct with two fields, `Username` and `Password`, both of which are strings.
//	登录必填信息
// The `binding:"required"` part is a validation rule. It means that the `Username` and `Password` fields are required.
//
// The `json:"username"` part is a tag. It means that the `Username` field will be serialized to JSON as `username`.
//
// The `json:"password"` part is a tag. It means that the `Password` field will be serialized to JSON as `password
// @property {string} Username - The username of the user.
// @property {string} Password - The password of the user.
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
