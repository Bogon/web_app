package post

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"webapp.io/controllers/requestHandler"
	"webapp.io/controllers/responseCode"
	"webapp.io/controllers/responseHandler"
	"webapp.io/logic/post"
	"webapp.io/models"
)

// CreatePostHandler is a function that takes a pointer to a gin.Context and returns nothing.
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数以及参数的校验
	p := new(models.Post)
	// 从 c 中获取当前登录用户ID
	userID, err := requestHandler.GetCurrentUserID(c)
	p.AuthorID = userID

	if err := c.ShouldBindJSON(p); err != nil {
		responseHandler.ResponseError(c, responseCode.CodeInvalidParam)
		return
	}

	if err != nil {
		responseHandler.ResponseError(c, responseCode.CodeToLogin)
		return
	}

	// 2. 创建帖子
	if err := post.CreatePost(p); err != nil {
		zap.L().Error(" post.CreatePost(p) failed", zap.Error(err))
		responseHandler.ResponseError(c, responseCode.CodeServerBasy)
		return
	}

	// 3. 返回响应
	responseHandler.ResponseSuccess(c, nil)
}

// GetPostDetailHandler is a function that takes a pointer to a gin.Context and returns nothing.
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取到postID参数并校验
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		return
	}
	// 2. 根据id 查询post 信息
	data, err := post.GetPostDetailById(id)
	if err != nil {
		zap.L().Error("post.GetPostDetail(id) failed", zap.Error(err))
		responseHandler.ResponseError(c, responseCode.CodeInvalidParam)
		return
	}
	// 3. 返回数据
	responseHandler.ResponseSuccess(c, data)

}

// GetPostListHandler gets a list of posts.
func GetPostListHandler(c *gin.Context) {

	var (
		page int64
		size int64
		err  error
	)

	page, size, err = requestHandler.GetPageInfo(c)
	if err != nil {
		page = 0
		size = 20
	}

	// 获取列表信息
	data, err := post.GetPostList(page, size)
	if err != nil {
		zap.L().Error("post.GetPostList() failed", zap.Error(err))
		return
	}
	// 返回数据
	responseHandler.ResponseSuccess(c, data)
}

// GetPostListSortedHandler gets a list of posts.
// 修改帖子列表接口，根据的用户传递过来的参数动态获取参数列表
// 按照 创建时间排序 或者 按照分数排序
// 1. 获取参数
// 2. redis 查询时间id 列表
// 3. 根据id到帖子数据库获取帖子详情信息
func GetPostListSortedHandler(c *gin.Context) {

	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	// 获取分页数据
	err := c.ShouldBindQuery(p)
	if err != nil {
		responseHandler.ResponseError(c, responseCode.CodeInvalidParam)
		return
	}
	// c.ShouldBindJSON() // 当请求中携带的参数是json时，使用该方法绑定参数
	// 获取列表信息
	//data, err := post.GetPostListSort(&p)
	data, err := post.GetPostListNew(p) // 更新接口，合二为一
	if err != nil {
		zap.L().Error("post.GetPostListSort() failed", zap.Error(err))
		responseHandler.ResponseError(c, responseCode.CodeServerBasy)
		return
	}
	// 返回数据
	responseHandler.ResponseSuccess(c, data)
}

//// GetCommunityPostListHandler 根据社区查询帖子列表
//// GetCommunityPostListHandler gets a list of posts for a community
//func GetCommunityPostListHandler(c *gin.Context) {
//	p := &models.ParamPostList{
//		Page:  1,
//		Size:  10,
//		Order: models.OrderTime,
//	}
//	// 获取分页数据
//	err := c.ShouldBindQuery(p)
//	if err != nil {
//		responseHandler.ResponseError(c, responseCode.CodeInvalidParam)
//		return
//	}
//	// c.ShouldBindJSON() // 当请求中携带的参数是json时，使用该方法绑定参数
//
//	// 获取列表信息
//	data, err := post.GetCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("post.GetCommunityPostList() failed", zap.Error(err))
//		responseHandler.ResponseError(c, responseCode.CodeServerBasy)
//		return
//	}
//	// 返回数据
//	responseHandler.ResponseSuccess(c, data)
//}
