package post

import (
	"go.uber.org/zap"
	daoPost "webapp.io/dao/mysql"
	"webapp.io/models"
	"webapp.io/pkg/snowflakeID"
)

// CreatePost creates a new post in the database.
func CreatePost(p *models.Post) (err error) {
	// 1. 生成 post id
	p.PostID = snowflakeID.GetId()
	// 2. 保存在数据库
	if err = daoPost.CreatePost(p); err != nil {
		zap.L().Error("post.CreatePost(p) failed", zap.Error(err))
		return
	}
	// 3. 返回
	return
}

// GetPostDetailById `GetPostDetail` returns a `*models.Post` and an `error` for a given `int64`
func GetPostDetailById(id int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合接口需要的数据
	post, err := daoPost.GetPostById(id)
	if err != nil {
		zap.L().Error(
			"daoPost.GetPostById(id) failed",
			zap.Int64("post_id", id),
			zap.Error(err))
		return
	}
	// 2. 根据ID查询作者信息
	user, err := daoPost.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error(
			"daoPost.GetUserById(post.AuthorID) failed",
			zap.Int64("authorID", post.AuthorID),
			zap.Error(err))
		return
	}

	// 3. 根据 community_id 获取 社区信息
	communityDetail, err := daoPost.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error(
			"daoPost.GetCommunityDetail(post.CommunityID) failed",
			zap.Int64("communityID", post.CommunityID),
			zap.Error(err))
		return
	}

	// 4. 组装数据返回
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: communityDetail,
	}
	return
}

// GetPostList > GetPostList returns a list of posts and an error
func GetPostList() (data []*models.ApiPostDetail, err error) {
	list, err := daoPost.GetPostList()
	if err != nil {
		zap.L().Error("daoPost.GetPostList() failed", zap.Error(err))
		return
	}

	data = make([]*models.ApiPostDetail, 0, len(list))

	for _, post := range list {
		// 1. 根据ID查询作者信息
		user, err := daoPost.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error(
				"daoPost.GetUserById(post.AuthorID) failed",
				zap.Int64("authorID", post.AuthorID),
				zap.Error(err))
			continue
		}

		// 2. 根据 community_id 获取 社区信息
		communityDetail, err := daoPost.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error(
				"daoPost.GetCommunityDetail(post.CommunityID) failed",
				zap.Int64("communityID", post.CommunityID),
				zap.Error(err))
			continue
		}

		// 3. 组装数据返回
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: communityDetail,
		}

		data = append(data, postDetail)
	}
	return
}
