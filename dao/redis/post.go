package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"webapp.io/models"
)

func getIdsFromKey(key string, page, size int64) ([]string, error) {
	// 2. 确定要查询索引起始点
	start := (page - 1) * p.Size
	end := start + size - 1
	// 3. 使用 ZRevRange 查询数据
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostIDsInOrder returns an error if it fails to get the post IDs in order
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从 redis 获取 ids
	// 1.根据用户请求中获取的order参数确定要查询的redis的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2. 确定要查询索引起始点
	return getIdsFromKey(key, p.Page, p.Size)
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

// GetCommunityPostIDsInOrder returns an error if it fails to get the post IDs in order
// 根据社区查询 ids
func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 使用 zinterscore 把分区的帖子 zset 与帖子分数的 zset 生成一个新的 zset
	// 针对新的 zset 按照之前的逻辑获取数据
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	// 社区key
	communityKey := getRedisKey(KeyCommunityZSetPrefix + strconv.Itoa(int(p.CommunityID)))
	// 利用缓存key减少 zinterscore 执行次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(key).Val() < 1 { // 不存在
		// 不存在需要计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, communityKey, orderKey)
		pipeline.Expire(key, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}

	// 存在的话直接根据key查询ids
	return getIdsFromKey(key, p.Page, p.Size)

}
