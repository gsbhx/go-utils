package redis

import (
	"fmt"
	redisconf "github.com/gsbhx/go-utils/config/cache/redis"
	redigo "github.com/garyburd/redigo/redis"
	"time"
)

var pool *redigo.Pool

func NewRedisClilent(r redisconf.RedisConfigure) {
	address := fmt.Sprintf("%s:%d", r.Host, r.Port)
	pool = &redigo.Pool{
		MaxIdle:     int(r.MaxidleNum),
		MaxActive:   int(r.MaxactiveNum),
		IdleTimeout: time.Second * 5,
		Wait:        true,
		Dial: func() (redigo.Conn, error) {
			conn, err := redigo.Dial(
				"tcp",
				address,
				redigo.DialPassword(r.Pwd),
				redigo.DialDatabase(0),
				redigo.DialConnectTimeout(5*time.Second),
				redigo.DialReadTimeout(5*time.Second),
				redigo.DialWriteTimeout(5*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}


