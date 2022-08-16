package responseHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp.io/controllers/responseCode"
)

/* 统一处理返回信息，设置返回的错误，定义返回样式
{
	"code": 10002,		// 程序中的错误码
	"msg": "message",	// 错误提示信息
	"data: {}			// 数据
}

*/

// ResponseData `ResponseData` is a struct with three fields, `Code`, `Msg`, and `Data`.
//
// The `Code` field is an `int` and is required.
//
// The `Msg` field is an `interface{}` and is required.
//
// The `Data` field is an `interface{}` and is required.
// @property {int} Code - The status code of the request.
// @property Msg - The message returned by the interface, which is usually used to return the error message.
// @property Data - The data returned by the server
type ResponseData struct {
	Code responseCode.ResCode `json:"code"`
	Msg  interface{}          `json:"msg"`
	Data interface{}          `json:"data"`
}

// ResponseError It returns an error response to the client.
func ResponseError(c *gin.Context, code responseCode.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		code,
		code.Msg(),
		nil,
	})
}
func ResponseWithMsg(c *gin.Context, code responseCode.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		code,
		msg,
		nil,
	})
}

// ResponseSuccess It returns a JSON response with the data passed in.
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		responseCode.CodeSuccess,
		responseCode.CodeSuccess.Msg(),
		data,
	})
}
