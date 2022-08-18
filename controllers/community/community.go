package community

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
