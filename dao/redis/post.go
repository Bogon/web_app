package redis

import (
	"github.com/go-redis/redis"
	"webapp.io/models"
)

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

// GetPostVoteData > This function will get the data from the database and return it as a struct
// 根据 ids 查询每篇帖子投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//for _, id := range ids {
	//data = make([]int64, len(ids))
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	// 查询key中分数是1的元素个数 -> 统计每篇帖子的赞成票数
	//	v := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}

	// 使用pipeline 一次发送多条命令，减少RTT
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")

	}
	exec, err := pipeline.Exec()
	if err != nil {
		return
	}
	data = make([]int64, 0, len(ids))
	for _, value := range exec {
		count := value.(*redis.IntCmd).Val()
		data = append(data, count)
	}
	return
}
