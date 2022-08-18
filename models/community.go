package models

import "time"

type Community struct {
	ID   int    `db:"community_id"`
	Name string `db:"community_name"`
}

type CommunityDetail struct {
	ID           int       `db:"community_id" json:"id"`
	Name         string    `db:"community_name" json:"name"`
	Introduction string    `db:"introduction" json:"introduction,omitempty"`
	CreateTime   time.Time `db:"create_time" json:"createTime"`
	UpdateTime   time.Time `db:"update_time" json:"updateTime"`
}
