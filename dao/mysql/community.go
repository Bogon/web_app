package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"webapp.io/models"
)

// GetCommunityList > GetCommunityList returns a slice of pointers to Community structs
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community;`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetail(id int64) (detail *models.CommunityDetail, err error) {
	detail = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time, update_time from community where community_id = ?;`
	if err = db.Get(detail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return detail, err
}
