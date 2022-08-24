package vote

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"webapp.io/controllers/requestHandler"
	"webapp.io/controllers/responseCode"
	"webapp.io/controllers/responseHandler"
	"webapp.io/controllers/validatorHandler"
	"webapp.io/logic/voted"
	"webapp.io/models"
)

// PostVoteHandler > PostVoteHandler is a function that takes a pointer to a gin.Context and returns nothing
// PostVoteHandler 为帖子投票
// @Summary 为帖子投票接口
// @Description 在允许的时间内可以发起投票并进行投票
// @Tags 投票相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamVoteData false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} responseHandler.ResponseData
// @Router /api/v1/vote [post]
func PostVoteHandler(c *gin.Context) {

	// 1. 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			responseHandler.ResponseError(c, responseCode.CodeInvalidParam)
			return
		}
		errData := validatorHandler.RemoveTopStruct(errs.Translate(validatorHandler.Trans)) // 翻译并去除错误信息中的结构体信息
		responseHandler.ResponseWithMsg(c, responseCode.CodeInvalidParam, errData)
		return
	}

	// 获取当前用户信息
	userId, err := requestHandler.GetCurrentUserID(c)
	if err != nil {
		responseHandler.ResponseError(c, responseCode.CodeToLogin)
	}

	if err := voted.VoteForPost(userId, p); err != nil {
		zap.L().Error("voted.VoteForPost failed", zap.Error(err))
		responseHandler.ResponseWithMsg(c, responseCode.CodeServerBasy, err.Error())
		return
	}
	responseHandler.ResponseSuccess(c, nil)
}
