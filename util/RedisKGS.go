package util

import (
	"user/db"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

type RedisKGS struct {
	redis     *db.RedisDb
	nextIdKey string
}

func (redisKGS *RedisKGS) NextId() int64 {
	c := redisKGS.redis.Connection()
	defer c.Close()
	nextId, _ := redis.Int64(c.Do("INCR", redisKGS.nextIdKey))
	//beego.Debug("NextId:", nextId)
	return nextId
}

//声明kgs
var (
	UserKGS *RedisKGS //user ID自增长
)

func init() {
	nameSpace := "kgs:" + beego.AppConfig.String("kgs.nameSpace")
	nextIdKey := nameSpace + ":nextId"

	UserKGS = &RedisKGS{
		redis:     db.UserRedis,
		nextIdKey: nextIdKey,
	}
	initRedisKGS(UserKGS)
	beego.Informational("init redis KGS nameSpace:", nameSpace)
}

func initRedisKGS(kgs *RedisKGS) {
	c := kgs.redis.Connection()
	defer c.Close()
	offset := beego.AppConfig.DefaultInt64("kgs.offset", 100000)
	is_key_exit, _ := redis.Bool(c.Do("EXISTS", kgs.nextIdKey))
	if !is_key_exit {
		_, err := c.Do("SET", kgs.nextIdKey, offset)
		if err != nil {
			beego.Error("redis KGS int error:", err)
		}
	}
	/*
		nextId, err := redis.Int64(c.Do("GET", kgs.nextIdKey))
		if err != nil {
			beego.Error("redis KGS int error:", err)
		} else if nextId < offset {
			_, err = c.Do("SET", kgs.nextIdKey, nextId)
			if err != nil {
				beego.Error("redis KGS int error:", err)
			}
		}
	*/
}
