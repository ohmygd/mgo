package redis

import (
	"errors"
	"time"
	"github.com/garyburd/redigo/redis"
)

type Locker struct {
	Key    string
	Error  error
}

// 获取锁, key尽量复杂, 才能减少出错的可能
func Lock(key string) (locker *Locker) {
	locker = &Locker{Key: key}

	redisConn := pool.Get()
	reply, _ := redis.String(redisConn.Do("SET", key, 1, "EX", 60, "NX"))
	redisConn.Close()

	if reply != "OK" {
		locker.Error = errors.New("locker failed.")
	}
	return
}

func (lock *Locker) Close() {
	if lock.Error == nil {
		redisConn := pool.Get()
		redisConn.Do("DEL", lock.Key)
		redisConn.Close()
	}
}

func TryLock(key string, timeout time.Duration) (locker *Locker) {
	locker = &Locker{Key: key}

	start := time.Now()
	for time.Now().Sub(start) < timeout {
		redisConn := pool.Get()
		// 锁定60s, 一个锁正常不会超过60s, 超过就自动删除
		reply, _ := redis.String(redisConn.Do("SET", key, 1, "EX", 60, "NX"))
		redisConn.Close()

		if reply == "OK" {
			return
		}

		time.Sleep(200 * time.Millisecond)
	}

	locker.Error = errors.New("locker timeout.\n")
	return
}