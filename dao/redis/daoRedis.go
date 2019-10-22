package redis

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/ohmygd/mgo/config"
	"github.com/ohmygd/mgo/merror"
	"github.com/ohmygd/mgo/pc"
	"time"
)

type DaoRedis struct {
	Conn redis.Conn
}

type GetType string

const (
	MaxIdleC              = 10
	MaxActiveC            = 0
	IdleTimeOutC          = 300
	GetTypeString GetType = "string"
	GetTypeInt            = "int"
	GetTypeBool           = "bool"
	GetTypeByte           = "byte"
)

var pool *redis.Pool

func init() {
	var maxIdleC, maxActiveC int
	var idleTimeOutC time.Duration

	host := config.GetRedisMsg("host")
	port := config.GetRedisMsg("port")
	maxIdle := config.GetRedisMsg("maxIdle")
	maxActive := config.GetRedisMsg("maxActive")
	idleTimeOut := config.GetRedisMsg("idleTimeOut")

	// 如果redis连接池信息为空, 设置初始信息
	if maxIdle != nil {
		maxIdleC = int(maxIdle.(float64))
	} else {
		maxIdleC = MaxIdleC
	}

	if maxActive != nil {
		maxActiveC = int(maxActive.(float64))
	} else {
		maxActiveC = MaxActiveC
	}

	if idleTimeOut != nil {
		idleTimeOutC = time.Duration(idleTimeOut.(float64))
	} else {
		idleTimeOutC = IdleTimeOutC
	}

	if host == "" || port == "" {
		panic("redis config lost")
	}

	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     maxIdleC,     //最初的连接数量
		MaxActive:   maxActiveC,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: idleTimeOutC, //连接关闭时间
		Wait:        true,
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			c, err := redis.Dial("tcp", host.(string)+":"+port.(string))
			if err != nil {
				err = merror.New(pc.ErrorRedisCon)
				return nil, err
			}

			// 授权验证
			password := config.GetRedisMsg("password")
			if password != nil {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					panic(fmt.Sprintf("redis auth error. err: %s", err))
				}
			}

			return c, err
		},
	}

}

func Pool() *redis.Pool {
	return pool
}

func (d *DaoRedis) BaseSet(c redis.Conn, key, value interface{}) (err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	_, err = c.Do("set", key, value)

	return
}

func (d *DaoRedis) BaseGet(c redis.Conn, key interface{}, getType GetType) (r interface{}, err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	switch getType {
	case GetTypeInt:
		return redis.Int(c.Do("get", key))
	case GetTypeBool:
		return redis.Bool(c.Do("get", key))
	case GetTypeString:
		return redis.String(c.Do("get", key))
	case GetTypeByte:
		return redis.Bytes(c.Do("get", key))
	}

	return nil, errors.New("getType error. not set")
}

func (d *DaoRedis) BaseSetEx(c redis.Conn, key interface{}, seconds int, value interface{}) (err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	_, err = c.Do("setex", key, seconds, value)

	return err
}

func (d *DaoRedis) BaseDel(c redis.Conn, key interface{}) (err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	_, err = c.Do("del", key)

	return err
}

func (d *DaoRedis) BaseMSet(c redis.Conn, p ...interface{}) (err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	_, err = c.Do("mset", p...)

	return
}

func (d *DaoRedis) BaseMGet(c redis.Conn, keys ...interface{}) (r []interface{}, err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	var res interface{}
	res, err = c.Do("mget", keys...)

	r = res.([]interface{})

	return
}

func (d *DaoRedis) Expire(c redis.Conn, key interface{}, second int) (err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	_, err = c.Do("expire", key, second)

	return
}

func (d *DaoRedis) DecrBy(c redis.Conn, key interface{}, by int) (r int, err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	var res interface{}
	res, err = c.Do("decrby", key, by)

	r = int(res.(int64))

	return
}

func (d *DaoRedis) IncrBy(c redis.Conn, key interface{}, by int) (r int, err error) {
	if c == nil {
		c = pool.Get()
		defer c.Close()
	}

	var res interface{}
	res, err = c.Do("incrby", key, by)

	r = int(res.(int64))

	return
}
