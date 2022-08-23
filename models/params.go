package models

// 定义请求参数结构体

const (
	OrderTime  = "time"  // 按照时间排序
	OrderScore = "score" // 按照分数排序
)

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

// ParamVoteData `ParamVoteData` is a struct with two fields, `PostID` and `Direction`.
//
// The `PostID` field is an `int64` and the `Direction` field is an `int`.
//
// The `json` tags on the fields tell the `json` package how to encode and decode the fields.
//
// The `json` package will ignore fields that are not exported (fields that start with a lowercase letter).
//
// The `json` package will also ignore fields that have a `json` tag with the value `"-"`.
// @property {int64} PostID - The ID of the post to be voted on
// @property {int} Direction - 1 is a positive vote, -1 is a negative vote
type ParamVoteData struct {
	// UserID 从当前请求中获取当前用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)取消投票(0); 只能是其中的一个
}

// ParamPostList `ParamPostList` is a struct with fields `Page`, `Size`, and `Order`.
//
// The `json:"page"` annotation tells the code generator to use the name `page` when serializing and deserializing the
// field.
//
// The `int64` type tells the code generator to use the `int64` type when serializing and deserializing the field.
//
// The `json:"page"` annotation tells the code generator to use the name `page` when serializing and deserializing the
// field.
//
// The `int64` type
// @property {int64} Page - The page number of the list
// @property {int64} Size - The number of records per page
// @property {string} Order - The order of the list, the default is "desc"
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
