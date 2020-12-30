package redis

type RedisInterface interface {
	Exists(key string) *Cbk
	Set(key string, value interface{}) *Cbk
	Get(key string) *Cbk
	Expire(key string, times int) *Cbk
	Del(key string) *Cbk
	Incr(key string) *Cbk
	Decr(key string) *Cbk
	HSet(key string, k, v interface{}) *Cbk
	HmSetHash(key string, kvs map[string]interface{}) *Cbk
	Hexists(key string, field string) *Cbk
	Hmget(key string, fields []string) *Cbk
	Hdel(key string, field string) *Cbk
	Hscan(key string, field string, count int) *Cbk

	//SET操作
	Sadd(key string, values []interface{}) *Cbk
	Sismember(key string, value string) *Cbk
	Srem(key string, value string) *Cbk

	//LIST操作
	Rpush(key string, value interface{}) *Cbk
	Lpush(key string, value interface{}) *Cbk
	Lpop(key string) *Cbk
	Rpop(key string) *Cbk
}
