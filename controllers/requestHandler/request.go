package requestHandler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"webapp.io/middlewares/jwtauth"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUserID Get the user ID from the context, if it's not there, return an error
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
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

// GetPageInfo > GetPageInfo gets the page and size from the request and returns them as integers
func GetPageInfo(c *gin.Context) (page, size int64, err error) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return
	}

	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		return
	}
	return
}
