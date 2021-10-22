package redis

import (
	"sort"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

type RedisService struct {
	address string
	auth    string
	pool    *redigo.Pool
}

func CreateRedisService(address string, auth string) (service *RedisService, err error) {
	service = &RedisService{
		address: address,
		auth:    auth,
	}
	err = service.init()
	return
}

func (service *RedisService) init() (err error) {
	pool := &redigo.Pool{
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   3, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", service.address)
			if err != nil {
				return nil, err
			}
			if service.auth != "" {
				if _, err := c.Do("auth", service.auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("ping")
			return err
		},
	}
	service.pool = pool
	return
}

func (service *RedisService) Keys(pattern string, size int) (count int, keys []string, err error) {

	client := service.pool.Get()
	defer client.Close()
	var reply interface{}
	reply, err = client.Do("keys", pattern)
	if err != nil {
		return
	}
	if reply != nil {
		var list []string
		list, err = redigo.Strings(reply, err)

		sor := sort.StringSlice(list)
		sor.Sort()

		count = len(list)
		if count <= size || size <= 0 {
			keys = list
		} else {
			for index, one := range list {
				if index < size {
					keys = append(keys, one)
				} else {
					break
				}
			}
		}
	}
	return
}

func (service *RedisService) KeyType(key string) (keyType string, err error) {

	client := service.pool.Get()
	defer client.Close()
	var reply interface{}
	reply, err = client.Do("type", key)
	if err != nil {
		return
	}
	if reply != nil {
		keyType, err = redigo.String(reply, err)
	}
	return
}

func (service *RedisService) Get(key string) (valueInfo ValueInfo, err error) {
	var keyType string
	keyType, err = service.KeyType(key)
	if err != nil {
		return
	}
	var value interface{}
	client := service.pool.Get()
	defer client.Close()
	if keyType == "none" {

	} else if keyType == "string" {
		var reply interface{}
		reply, err = client.Do("get", key)
		if err != nil {
			return
		}
		if reply != nil {
			value, err = redigo.String(reply, err)
		}
	} else if keyType == "list" {
		var reply interface{}
		reply, err = client.Do("llen", key)
		if err != nil {
			return
		}
		var len int64
		len, err = redigo.Int64(reply, err)
		if err != nil {
			return
		}
		reply, err = client.Do("lrange", key, 0, len)
		if err != nil {
			return
		}
		value, err = redigo.Strings(reply, err)
	} else if keyType == "set" {
		var reply interface{}
		reply, err = client.Do("smembers", key)
		if err != nil {
			return
		}
		value, err = redigo.Strings(reply, err)
	} else if keyType == "hash" {
		var reply interface{}
		reply, err = client.Do("hgetall", key)
		if err != nil {
			return
		}
		value, err = redigo.StringMap(reply, err)
	} else {
		println(keyType)
	}
	valueInfo.Type = keyType
	valueInfo.Value = value

	return
}

func (service *RedisService) Set(key string, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("set", key, value)
	return
}

func (service *RedisService) Sadd(key string, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("sadd", key, value)
	return
}

func (service *RedisService) Srem(key string, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("srem", key, value)
	return
}

func (service *RedisService) Lpush(key string, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("lpush", key, value)
	return
}

func (service *RedisService) Rpush(key string, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("rpush", key, value)
	return
}

func (service *RedisService) Lset(key string, index int64, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("lset", key, index, value)
	return
}

func (service *RedisService) Lrem(key string, count int64, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("lrem", key, count, value)
	return
}

func (service *RedisService) Hset(key string, field string, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("hset", key, field, value)
	return
}

func (service *RedisService) Hdel(key string, field string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("hdel", key, field)
	return
}

func (service *RedisService) Del(key string) (count int, err error) {
	count = 0
	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("del", key)
	if err == nil {
		count++
	}
	return
}

func (service *RedisService) DelPattern(pattern string) (count int, err error) {
	count = 0
	var list []string
	_, list, err = service.Keys(pattern, 0)

	client := service.pool.Get()
	defer client.Close()

	for _, key := range list {
		_, err = client.Do("del", key)
		if err == nil {
			count++
		} else {
			return
		}
	}
	return
}
