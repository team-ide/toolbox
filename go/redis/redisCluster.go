package redis

import (
	"context"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClusterService struct {
	address      string
	auth         string
	redisCluster *redis.ClusterClient
}

func CreateRedisClusterService(address string, auth string) (service *RedisClusterService, err error) {
	service = &RedisClusterService{
		address: address,
		auth:    auth,
	}
	err = service.init()
	return
}

func (service *RedisClusterService) init() (err error) {
	redisCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        strings.Split(service.address, ";"),
		DialTimeout:  100 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
		Password:     service.auth,
	})
	service.redisCluster = redisCluster
	return
}

func (service *RedisClusterService) Keys(pattern string, size int) (count int, keys []string, err error) {
	cmd := service.redisCluster.Keys(context.TODO(), pattern)
	var list []string
	list, err = cmd.Result()
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
func (service *RedisClusterService) Get(key string) (value string, err error) {

	cmd := service.redisCluster.Get(context.TODO(), key)
	value, err = cmd.Result()
	return
}

func (service *RedisClusterService) Set(key string, value string) (err error) {
	cmd := service.redisCluster.Set(context.TODO(), key, value, time.Duration(0))

	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Del(key string) (count int, err error) {
	count = 0
	cmd := service.redisCluster.Del(context.TODO(), key)
	_, err = cmd.Result()
	if err == nil {
		count++
	}
	return
}

func (service *RedisClusterService) DelPattern(pattern string) (count int, err error) {
	count = 0
	cmd := service.redisCluster.Keys(context.TODO(), pattern)
	var list []string
	list, err = cmd.Result()
	for _, key := range list {
		var num int
		num, err = service.Del(key)
		if err != nil {
			return
		}
		count += num
	}
	return
}
