package db

import (
	"github.com/go-redis/redis"
	"github.com/zngue/go_tool/src/sign_chan"
	"time"
)
func RedisConnet()  {
	RedisSorce(Config.Redis.DBNum)
}
func RedisSorce(DBNum int)  {
	redisC :=Config.Redis
	RedisConn = redis.NewClient(&redis.Options{
		Addr:         redisC.Host + ":"+redisC.Port,
		Password:     redisC.Password,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     30,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: 10,
		DB:           DBNum,
	})
	pong, err := RedisConn.Ping().Result()
	if err != nil {
		sign_chan.SignLog("redis:错误",pong,err)
	}
	
}
func Redis(dbNum int) *redis.Client {
	redisC :=Config.Redis
	RedisCsnn := redis.NewClient(&redis.Options{
		Addr:         redisC.Host + ":"+redisC.Port,
		Password:     redisC.Password,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     30,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: 10,
		DB:           dbNum,
	})
	pong, err := RedisCsnn.Ping().Result()
	if err != nil {
		sign_chan.SignLog("redis:错误",pong,err)
	}
	return  RedisCsnn
}
