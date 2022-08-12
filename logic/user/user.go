package user

import (
	"webapp.io/dao/mysql/user"
	"webapp.io/pkg/snowflakeID"
)

// 存放业务逻辑的地方

// SignUp `SignUp()` is a function that allows a user to sign up for an account
func SignUp() {
	// 1. 判断用户是否存在
	user.QueryUserByUsername()

	// 2. 生成 UID
	snowflakeID.GetId()

	// 3. 密码加密

	// 4. 保存到数据库
	user.InsertUser()
}
