package voted

import (
	"go.uber.org/zap"
	"strconv"
	"webapp.io/dao/redis"
	"webapp.io/models"
)

// 推荐阅读
//	基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目的简化版本的投票分数
// 1. 用户投一票+432  86400/200 200张赞成票可以给你的帖子续一天  《Redis实战》

// 投票的集中情况
/*
a. direction 为 1 时, 有两种情况：
	1. 之前没有投过票，现在投赞成票		----> 更新分数和投票记录
	2. 之前投反对票，现在改投赞成票		----> 更新分数和投票记录

b. direction 为 0 时, 有两种情况：
	1. 之前投过赞成票，现在取消投票		----> 更新分数和投票记录
	2. 之前投过返回票，现在取消投票		----> 更新分数和投票记录

c. direction 为 -1 时, 有两种情况：
	1. 之前没有投过票，现在投反对票		----> 更新分数和投票记录
	2. 之前投赞成票，现在改投反对票		----> 更新分数和投票记录

投票的限制：
1.每个帖子自发表之日起，在一周之内允许用户投票，超过一个星期不允许投票。
	1. 到期之后将redis中的赞成票票数和反对票数存储到mysql数据库中
	2. 到期之后删除保存key：KeyPostVotedZSetPrefix
*/

// 投票功能
// 1. 用户投票的数据
// 2.

// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	// 1. 投票限制
	// 2. 更新投票帖子分数
	// 3. 记录用户为该帖子投过票的数据
	zap.L().Debug("redis.VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction),
	)
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
