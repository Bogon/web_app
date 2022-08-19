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
