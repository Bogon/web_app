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
	// 获取分页参数
	pageStr := c.Query("page") // 分页
	sizeStr := c.Query("size") // 页面个数

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 0
	}

	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
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
