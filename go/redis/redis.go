package redis

import (
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
				if _, err := c.Do("AUTH", service.auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	service.pool = pool
	return
}

func (service *RedisService) Keys(pattern string, size int) (count int, keys []string, err error) {

	client := service.pool.Get()
	defer client.Close()
	var list []string
	list, err = redigo.Strings(client.Do("KEYS", pattern))
	count = len(list)
	if count <= size {
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
	return
}
func (service *RedisService) Get(key string) (value string, err error) {

	client := service.pool.Get()
	defer client.Close()
	value, err = redigo.String(client.Do("GET", key))
	return
}

func (service *RedisService) Set(key string, value string) (err error) {

	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("SET", key, value)
	return
}

func (service *RedisService) Del(key string) (count int, err error) {
	count = 0
	client := service.pool.Get()
	defer client.Close()
	_, err = client.Do("DEL", key)
	if err == nil {
		count++
	}
	return
}

func (service *RedisService) DelPattern(pattern string) (count int, err error) {
	count = 0
	client := service.pool.Get()
	defer client.Close()
	var list []string
	list, err = redigo.Strings(client.Do("KEYS", pattern))

	for _, key := range list {
		_, err = client.Do("DEL", key)
		if err == nil {
			count++
		} else {
			return
		}
	}
	return
}
