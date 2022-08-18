package community

import (
	"webapp.io/dao/mysql"
	"webapp.io/models"
)

// GetCommunityList returns a list of all the communities in the database
func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的 community 返回
	return mysql.GetCommunityList()
}

// GetCommunityDetail returns a CommunityDetail struct and an error
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetail(id)
}
