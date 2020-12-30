package redis

type Redis struct{}

//常规操作
func (s Redis) Exists(key string) *Cbk {
	c := pool.Get()
	defer c.Close()
	return GetCbk(c.Do("EXISTS", key))
}
func (s Redis) Set(key string, value interface{}) *Cbk {
	c := pool.Get()
	defer c.Close()
	return GetCbk(c.Do("SET", key, value))
}
func (s Redis) Get(key string) *Cbk {
	c := pool.Get()
	defer c.Close()
	return GetCbk(c.Do("GET", key))
}
func (s Redis) Expire(key string, times int) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("expire", key, times))
}
func (s Redis) Del(key string) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("DEL", key))
}
func (s Redis) Incr(key string) *Cbk {
	c := pool.Get()
	defer c.Close()
	return GetCbk(c.Do("INCR", key))
}
func (s Redis) Decr(key string) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("DECR", key))
}
func (s Redis) HSet(key string, k, v interface{}) *Cbk {
	c := pool.Get()
	defer c.Close()
	return GetCbk(c.Do("HSET", key, k, v))
}
func (s Redis) HmSetHash(key string, kvs map[string]interface{}) *Cbk {
	c := pool.Get()
	defer c.Close()
	args := []interface{}{key}
	for key, value := range kvs {
		args = append(args, key, value)
	}
	return GetCbk(c.Do("HMSET", args...))
}
func (s Redis) Hexists(key string, field string) *Cbk {
	c := pool.Get()
	defer c.Close()
	return GetCbk(c.Do("HEXISTS", key, field))
}
func (s Redis) Hmget(key string, fields []string) *Cbk {
	c := pool.Get()
	defer c.Close()
	args := []interface{}{key}
	for _, value := range fields {
		args = append(args, value)
	}
	return GetCbk(c.Do("HMGET", args...))
}
func (s Redis) Hdel(key string, field string) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("HDEL", key, field))
}
func (s Redis) Hscan(key string, field string, count int) *Cbk {
	c := pool.Get()
	defer c.Close()
	params := []interface{}{key, 0, "MATCH", field}
	if count > 0 {
		params = append(params, "count", count)
	}
	return GetCbk(c.Do("HSCAN", params...))
}

//SET操作
func (s Redis) Sadd(key string, values []interface{}) *Cbk {
	client := pool.Get()
	defer client.Close()
	var args = []interface{}{key}
	for _, value := range values {
		args = append(args, value)
	}
	return GetCbk(client.Do("SADD", args...))
}
func (s Redis) Sismember(key string, value string) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("SISMEMBER", key, value))
}
func (s Redis) Srem(key string, value string) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("SREM", key, value))
}

//LIST操作
func (s Redis) Rpush(key string, value interface{}) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("RPUSH", key, value))
}
func (s Redis) Lpush(key string, value interface{}) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("LPUSH", key, value))
}
func (s Redis) Lpop(key string) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("LPOP", key))
}
func (s Redis) Rpop(key string) *Cbk {
	client := pool.Get()
	defer client.Close()
	return GetCbk(client.Do("RPOP", key))
}

