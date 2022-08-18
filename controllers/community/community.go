package community

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"webapp.io/controllers/responseCode"
	"webapp.io/controllers/responseHandler"
	"webapp.io/logic/community"
)

//  ---- 跟社区相关的 ----

// GetCommunityHandler Get all the communities and return them in a list
func GetCommunityHandler(c *gin.Context) {
	// 查询到所有的社区，然后一列表的形式给出来(community_id, community_name)
	data, err := community.GetCommunityList()
	if err != nil {
		zap.L().Error("community.GetCommunityList() failed ", zap.Error(err))
		responseHandler.ResponseError(c, responseCode.CodeServerBasy) // 服务端错误不要暴露到外部
		return
	}
	responseHandler.ResponseSuccess(c, data)
}

// GetCommunityDetailHandler is a function that returns a gin.HandlerFunc
func GetCommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区ID
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		responseHandler.ResponseError(c, responseCode.CodeInvalidParam)
		return
	}
	// 根据 id 查询 社区分类详情
	data, err := community.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("community.GetCommunityList() failed ", zap.Error(err))
		responseHandler.ResponseWithMsg(c, responseCode.CodeInvalidParam, err.Error()) // 服务端错误不要暴露到外部
		return
	}
	responseHandler.ResponseSuccess(c, data)
}
