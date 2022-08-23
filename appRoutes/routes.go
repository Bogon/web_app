package appRoutes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp.io/controllers/community"
	"webapp.io/controllers/post"
	"webapp.io/controllers/userHanlder"
	"webapp.io/controllers/vote"
	"webapp.io/logger"
	"webapp.io/middlewares/jwtauth"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin 设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("api/v1")

	// 注册业务路由 - 注册
	v1.POST("signup", userHanlder.UserSignUpHandler)
	// 注册业务路由 - 登录
	v1.POST("login", userHanlder.UserLoginHandler)

	v1.Use(jwtauth.JWTAuthMiddleware())

	{
		// 获取社区分类
		v1.GET("community", community.GetCommunityHandler)
		v1.GET("community/:id", community.GetCommunityDetailHandler)

		// 发布帖子
		v1.POST("post", post.CreatePostHandler)
		v1.GET("post/:id", post.GetPostDetailHandler)
		v1.GET("/posts", post.GetPostListHandler)
		// 根据时间/分数获取帖子列表
		v1.GET("/postssorted", post.GetPostListSortedHandler)

		// 投票
		v1.POST("vote", vote.PostVoteHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
