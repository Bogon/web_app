package mysql

import "webapp.io/models"

// CreatePost > CreatePost creates a new post in the database
func CreatePost(p *models.Post) (err error) {
	// 执行 sql 语句入库
	sqlStr := `insert into post(post_id, title, content, author_id, community_id, status) values(?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.CommunityID, p.Status)
	return
}
