package models

import "time"

// Community is a struct with two fields, ID and Name, both of which are exported.
// @property {int} ID - The ID of the community.
// @property {string} Name - The name of the property.
type Community struct {
	ID   int64  `db:"community_id"`
	Name string `db:"community_name"`
}

// CommunityDetail `CommunityDetail` is a struct with fields `ID`, `Name`, `Introduction`, `CreateTime`, and `UpdateTime`.
//
// The `db` tags are used by the `sqlx` package to map the fields to the columns in the database.
//
// The `json` tags are used by the `json` package to map the fields to the keys in the JSON.
//
// The `omitempty` option tells the `json` package to omit the field if it is empty.
//
// The `time.Time` type is a struct with fields `
// @property {int} ID - The ID of the community.
// @property {string} Name - The name of the property in the struct.
// @property {string} Introduction - community introduction
// @property CreateTime - The time when the community was created.
// @property UpdateTime - The time when the community was last updated.
type CommunityDetail struct {
	ID           int64     `db:"community_id" json:"id,string"`
	Name         string    `db:"community_name" json:"name"`
	Introduction string    `db:"introduction" json:"introduction,omitempty"`
	CreateTime   time.Time `db:"create_time" json:"createTime"`
	UpdateTime   time.Time `db:"update_time" json:"updateTime"`
}
