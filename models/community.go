package models

type Community struct {
	ID           int    `db:"community_id"`
	Name         string `db:"community_name"`
	Introduction string `db:"introduction"`
}
