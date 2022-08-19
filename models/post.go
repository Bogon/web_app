package models

import "time"

// Post `Post` is a struct with fields `ID`, `PostID`, `Title`, `Content`, `AuthorID`, `CommunityID`, `Status`, `CreateTime`,
// and `UpdateTime`.
//
// The `db` tags are used by the `sqlx` package to map the struct fields to the database columns.
//
// The `json` tags are used by the `json` package to map the struct fields to the JSON keys.
//
// The `time.Time` type is used to represent a date and time.
//
// The `int64` type is used
// @property {int64} ID - The ID of the post.
// @property {int64} PostID - The ID of the post.
// @property {string} Title - The title of the post.
// @property {string} Content - The content of the post.
// @property {string} AuthorID - The ID of the user who created the post.
// @property {string} CommunityID - The ID of the community that the post belongs to.
// @property {int} Status - 0: normal, 1: deleted, 2: banned
// @property CreateTime - The time when the post was created.
// @property UpdateTime - The time when the post was last updated.
type Post struct {
	ID          int64     `db:"id" json:"id"`
	PostID      int64     `db:"post_id" json:"postID"`
	AuthorID    int64     `db:"author_id" json:"authorID" binding:"required"`
	CommunityID int64     `db:"community_id" json:"communityID" binding:"required"`
	Status      int32     `db:"status" json:"status"`
	Title       string    `db:"title" json:"title" binding:"required"`
	Content     string    `db:"content" json:"content" binding:"required"`
	CreateTime  time.Time `db:"create_time" json:"createTime"`
	UpdateTime  time.Time `db:"update_time" json:"updateTime"`
}
