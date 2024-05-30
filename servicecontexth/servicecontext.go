package servicecontexth

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/nsqio/go-nsq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	RedisClient *redis.Client
	Db          *gorm.DB

	RocketProducer     rocketmq.Producer
	RocketPushConsumer rocketmq.PushConsumer

	NsqProducer *nsq.Producer
}
