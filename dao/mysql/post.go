package mysql

import (
	"database/sql"
	"webapp.io/models"
)

// CreatePost > CreatePost creates a new post in the database
func CreatePost(p *models.Post) (err error) {
	// 执行 sql 语句入库
	sqlStr := `insert into post(post_id, title, content, author_id, community_id, status) values(?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.CommunityID, p.Status)
	return
}

// GetPostById It returns a post with the given id, or an error if the post doesn't exist
func GetPostById(id int64) (detail *models.Post, err error) {
	detail = new(models.Post)
	// 执行sql
	sqlStr := `select id, post_id, author_id, community_id, status, title, content, create_time, update_time from post where post_id = ?`
	// 判断 sql 执行完成之后，数据为空的情况
	if err = db.Get(detail, sqlStr, id); err != nil {
		// 如果 查询失败，则返回错误
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return detail, err
}

// GetPostList A function that returns a slice of pointers to Post structs and an error.
func GetPostList(page, size int64) (data []*models.Post, err error) {
	sqlStr := `select 
    id, post_id, author_id, community_id, status, title, content, create_time, update_time 
	from post
	limit ?,?
	`
	data = make([]*models.Post, 0, 2)
	err = db.Select(&data, sqlStr, (page-1)*size, size)
	return
}
