package mysql

import (
	"testing"
	"webapp.io/models"
	"webapp.io/settings"
)

func init() {
	dbCfg := &settings.MySQLConf{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "123456",
		Dbname:       "web_app",
		Port:         3306,
		OpenMaxConns: 200,
		IdelMaxConns: 50,
	}

	err := Init(dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	/*
		ID          int64     `db:"id" json:"id,string"`
			PostID      int64     `db:"post_id" json:"postID,string"`
			AuthorID    int64     `db:"author_id" json:"authorID,string" binding:"required"`
			CommunityID int64     `db:"community_id" json:"communityID,string" binding:"required"`
			Status      int32     `db:"status" json:"status"`
			Title       string    `db:"title" json:"title" binding:"required"`
			Content     string    `db:"content" json:"content" binding:"required"`
			CreateTime  time.Time `db:"create_time" json:"createTime"`
			UpdateTime  time.Time `db:"update_time" json:"updateTime"`
	*/
	p := &models.Post{
		ID:          198237912378192,
		PostID:      12,
		AuthorID:    12903810293,
		CommunityID: 2,
		Status:      1,
		Title:       " 这是一个测试数据",
		Content:     "黑雅黑",
	}

	err := CreatePost(p)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err: %v  \n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
