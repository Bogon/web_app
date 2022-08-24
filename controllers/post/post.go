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
// CreatePostHandler 发布帖子接口
// @Summary 发布帖子接口
// @Description 写入必要的信息提交新帖子保存到数据库
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.Post false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} nil
// @Router /api/v1/post [post]
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
// GetPostDetailHandler 查询帖子详情接口
// @Summary 查询帖子详情接口
// @Description 从数据库中查询帖子数据并返回帖子信息
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ApiPostDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models.ApiPostDetail
// @Router /api/v1/post [post]
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
// GetPostListHandler 查询帖子列表接口
// @Summary 查询帖子列表接口
// @Description 查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ApiPostDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} []models.ApiPostDetail
// @Router /api/v1/posts [post]
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

// GetPostListSortedHandler 升级版帖子列表接口 gets a list of posts.
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} []models.ApiPostDetail
// @Router /api/v1/postssorted [get]
func GetPostListSortedHandler(c *gin.Context) {
	// 修改帖子列表接口，根据的用户传递过来的参数动态获取参数列表
	// 按照 创建时间排序 或者 按照分数排序
	// 1. 获取参数
	// 2. redis 查询时间id 列表
	// 3. 根据id到帖子数据库获取帖子详情信息
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
