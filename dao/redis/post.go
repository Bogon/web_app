package redis

import "webapp.io/models"

// GetPostIDsInOrder returns an error if it fails to get the post IDs in order
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从 redis 获取 ids
	// 1.根据用户请求中获取的order参数确定要查询的redis的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2. 确定要查询索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3. 使用 ZRevRange 查询数据
	return rdb.ZRevRange(key, start, end).Result()
}
