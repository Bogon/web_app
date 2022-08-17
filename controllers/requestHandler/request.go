package requestHandler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"webapp.io/middlewares/jwtauth"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser Get the user ID from the context, if it's not there, return an error
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(jwtauth.CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
