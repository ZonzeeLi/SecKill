/**
    @Author:     ZonzeeLi
    @Project:    sk_layer
    @CreateDate: 1/10/2022
    @UpdateDate: xxx
    @Note:       redis初始化
**/

package setting

import (
	"github.com/go-redis/redis"
	"log"
	"sk_layer/config"
)

func InitRedis(host string, passWord string, db int, proxy2layerQueueName, layer2proxyQueueName string) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: passWord,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Printf("Connect redis failed. Error : %v", err)
	}

	config.SecLayerCtx.RedisConf = &config.RedisConf{
		RedisConn:            client,
		Proxy2layerQueueName: proxy2layerQueueName,
		Layer2proxyQueueName: layer2proxyQueueName,
	}
}
