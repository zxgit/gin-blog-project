/*
@Time : 2021/3/1 21:44
@Author : zhangxin
*/
package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimiterIface interface {
	//获取对应限流器的键值对名称
	Key(c *gin.Context) string
	//获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	//新增多个令牌桶
	AddBucket(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        //自定义键值对名称
	FillInterval time.Duration //间隔多久放n令牌桶
	Capacity     int64         //令牌桶的容量
	Quantum      int64         //每次达到间隔时间后所放的具体令牌数量
}
