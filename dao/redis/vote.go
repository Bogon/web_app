package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

// 一周之内可以发起投票

const (
	oneWeekSeconds = 7 * 24 * 3600 // 一周时间秒数
	scorePerVote   = 432           // 每一票所占的分数
)

var (
	ErrorVoteTimeExpire = errors.New("投票时间已过")
)

// 投票的集中情况：
/*
a. direction 为 1 时, 有两种情况：
	1. 之前没有投过票，现在投赞成票		----> 更新分数和投票记录  差值的绝对值：1	+432
	2. 之前投反对票，现在改投赞成票		----> 更新分数和投票记录  差值的绝对值：2	+432*2

b. direction 为 0 时, 有两种情况：
	1. 之前投过赞成票，现在取消投票		----> 更新分数和投票记录  差值的绝对值：1	-432
	2. 之前投过返回票，现在取消投票		----> 更新分数和投票记录  差值的绝对值：1	+432

c. direction 为 -1 时, 有两种情况：
	1. 之前没有投过票，现在投反对票		----> 更新分数和投票记录  差值的绝对值：1	-432
	2. 之前投赞成票，现在改投反对票		----> 更新分数和投票记录  差值的绝对值：2	-432*2

投票的限制：
1.每个帖子自发表之日起，在一周之内允许用户投票，超过一个星期不允许投票。
	1. 到期之后将redis中的赞成票票数和反对票数存储到mysql数据库中
	2. 到期之后删除保存key：KeyPostVotedZSetPrefix
*/

// VoteForPost > VoteForPost() is a function that allows a user to vote for a post
func VoteForPost(userID, postID string, value float64) (err error) {
	// 1. 投票限制
	// 从redis中获取帖子时间：KeyPostTimeZSet
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	// 超过一周不可以发起投票
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrorVoteTimeExpire
	}

	// 2/3 需要放在一个pipeline中进行操作
	pipeline := rdb.TxPipeline()
	// 2. 更新投票帖子分数
	// 先查当前用户给当前帖子的投票记录
	oValue := rdb.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()

	var op float64 // 计算差值
	if value > oValue {
		op = 1
	} else {
		op = -1
	}

	diff := math.Abs(oValue - value) // 计算两次投票的差值
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)

	// 3. 记录用户为该帖子投过票的数据
	if value == 0 {
		// 移除 key
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID})
	}
	_, err = pipeline.Exec()
	return
}

// CreatePost > This function creates a new post in the database
func CreatePost(pID int64) (err error) {

	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pID,
	})
	_, err = pipeline.Exec()
	return
}
