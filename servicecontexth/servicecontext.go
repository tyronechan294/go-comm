package servicecontexth

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	RedisClient *redis.Client
	Db          *gorm.DB

	RocketProducer     rocketmq.Producer
	RocketPushConsumer rocketmq.PushConsumer

	TronGrpcClient *client.GrpcClient
}
