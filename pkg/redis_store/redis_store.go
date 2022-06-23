package redis_store

import (
	"context"
	"easy-admin-go-service/pkg/viper"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	RedisDb *redis.Client
)

func init() {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     viper.SrvConf.Redis.Host,
		Password: viper.SrvConf.Redis.PassWord, // no password set
		DB:       viper.SrvConf.Redis.Db,       // use default DB
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		log.Printf("The connection pool connection in the redis no sql is fail", err.Error())
	}
	log.Printf("The connection pool connection in the redis no sql is successful")
}
