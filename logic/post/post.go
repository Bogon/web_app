package post

import (
	"go.uber.org/zap"
	daoPost "webapp.io/dao/mysql"
	"webapp.io/dao/redis"
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

	if err = redis.CreatePost(p.PostID); err != nil {
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

// GetPostList > GetPostList() returns a list of post details, including the author's name, the post's content, and the community's
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	list, err := daoPost.GetPostList(page, size)
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

// GetPostListSort > GetPostListSort() is a function that gets the post list and sorts it by the author's name and community name
// 2. redis 查询时间id 列表
// 3. 根据id到帖子数据库获取帖子详情信息
func GetPostListSort(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {

	// 2. redis 查询时间id 列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}

	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0")
		return
	}
	// 3. 根据id到帖子MySQL数据库获取帖子详情信息

	list, err := daoPost.GetPostListSort(ids)
	if err != nil {
		zap.L().Error("daoPost.GetPostList() failed", zap.Error(err))
		return
	}

	data = make([]*models.ApiPostDetail, 0, len(ids))

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
