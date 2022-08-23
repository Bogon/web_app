package redis

import "fmt"

// redis key

// redis 中的key 尽量使用命名空间的方式，方便业务查询、拆分

const (
	KeyPrefix              = "webapp.io:"
	KeyPostTimeZSet        = "post:time"  // ZSet; 帖子以发帖时间作为分数
	KeyPostScoreZSet       = "post:score" // ZSet; 帖子以投票分数作为分数
	KeyPostVotedZSetPrefix = "post:voted" // 记录用户以及投票的类型；参数是帖子id：post_id
)

// It takes a string and returns a string
func getRedisKey(value string) string {
	return fmt.Sprintf("%v%v", KeyPrefix, value)
}
