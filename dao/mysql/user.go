package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
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
		return ErrorUserExist
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

// encryptPassword It takes a string, appends a secret string to it, and then returns the MD5 hash of the result
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Login takes a pointer to a ParamLogin struct and does nothing
func Login(p *models.User) (err error) {
	oPassword := p.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(p, sqlStr, p.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password == p.Password {
		return
	} else {
		return ErrorUserInvalidPassword
	}
}

// GetUserById > GetUserById returns a user by id
func GetUserById(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?;`
	if err = db.Get(user, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			return
		}
	}
	return
}
