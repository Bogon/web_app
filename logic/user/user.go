package user

import (
	"webapp.io/dao/mysql"
	"webapp.io/models"
	"webapp.io/pkg/snowflakeID"
)

// 存放业务逻辑的地方

// SignUp `SignUp()` is a function that allows a user to sign up for an account
func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2. 生成 UID
	userID := snowflakeID.GetId()

	// 构造一个user实例
	user := &models.User{UserID: userID, Username: p.Username, Password: p.Password}

	// 4. 保存到数据库
	return mysql.InsertUser(user)
}

// Login `Login` takes a `*models.ParamLogin` and returns an `error`
func Login(p *models.ParamLogin) error {
	// 判断用户是否存在
	u := &models.User{Username: p.Username, Password: p.Password}
	return mysql.Login(u)
}
