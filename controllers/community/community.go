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
// GetCommunityHandler 获取社区列表
// @Summary 发布帖子接口
// @Description 查询发布帖子社区列表
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} []models.Community
// @Router /api/v1/community [get]
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
// GetCommunityDetailHandler 获取社区详情
// @Summary 获取社区详情接口
// @Description 查询社区详情
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int64  false  "Community ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.CommunityDetail
// @Router /api/v1/community/:id [get]
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
