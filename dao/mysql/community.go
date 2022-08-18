package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"webapp.io/models"
)

// GetCommunityList > GetCommunityList returns a slice of pointers to Community structs
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name, introduction from community;`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
			return
		}
	}
	return
}
