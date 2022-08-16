package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"webapp.io/models"
)

const secret = "webapp.io"

// 把每一步数据库操作封装成一个函数
// 待 logic 层根据业务需求调用

// CheckUserExist queries a user by username
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser `SignUp()` is a function that allows a user to sign up for an account
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行 sql 语句入库
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
