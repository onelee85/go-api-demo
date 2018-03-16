package db

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

type RedisDb struct {
	redisPool *redis.Pool
}

func (redisDb *RedisDb) Connection() redis.Conn {
	return redisDb.redisPool.Get()
}

var (
	UserRedis *RedisDb
)

func init() {
	UserRedis = &RedisDb{
		redisPool: initRedisPool("userredis"),
	}
	beego.Informational("init redis MaxActive:", UserRedis.redisPool.MaxActive)

}

func initRedisPool(redisName string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     beego.AppConfig.DefaultInt(redisName+".maxidle", 10),   //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   beego.AppConfig.DefaultInt(redisName+".maxactive", 10), //最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: 180 * time.Second,                                      //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String(redisName+".host"))
			if err != nil {
				beego.Error("redis Dial error:", err)
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

/*
func redisTest() {
	c := UserRedis.Connection()
	defer c.Close()
	startTime := time.Now()
	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	fmt.Printf("用时：%s", time.Now().Sub(startTime))
}

func main() {
	for i := 0; i < 100; i++ {
		redisTest()
	}
}
*/
