package redisconf

var RedisConf *RedisConfigure
type RedisConfigure struct {
	Host string
	Port int
	MaxidleNum int
	MaxactiveNum int
	Pwd string
}