package community

import (
	"webapp.io/dao/mysql"
	"webapp.io/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的 community 返回
	return mysql.GetCommunityList()
}
