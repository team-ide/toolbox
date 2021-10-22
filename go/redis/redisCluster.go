package redis

import (
	"context"
	"sort"
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
	var list []string
	service.redisCluster.ForEachMaster(context.TODO(), func(ctx context.Context, client *redis.Client) (err error) {
		cmd := client.Keys(ctx, pattern)
		var ls []string
		ls, err = cmd.Result()
		if err != nil {
			return
		}
		list = append(list, ls...)
		return
	})
	sor := sort.StringSlice(list)
	sor.Sort()
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

func (service *RedisClusterService) KeyType(key string) (keyType string, err error) {
	cmd := service.redisCluster.Type(context.TODO(), key)
	keyType, err = cmd.Result()
	return
}

func (service *RedisClusterService) Get(key string) (valueInfo ValueInfo, err error) {
	var keyType string
	keyType, err = service.KeyType(key)
	if err != nil {
		return
	}
	var value interface{}

	if keyType == "none" {

	} else if keyType == "string" {
		cmd := service.redisCluster.Get(context.TODO(), key)
		value, err = cmd.Result()
	} else if keyType == "list" {

		cmd := service.redisCluster.LLen(context.TODO(), key)

		var len int64
		len, err = cmd.Result()
		if err != nil {
			return
		}
		cmdRange := service.redisCluster.LRange(context.TODO(), key, 0, len)
		value, err = cmdRange.Result()
	} else if keyType == "set" {
		cmd := service.redisCluster.SMembers(context.TODO(), key)
		value, err = cmd.Result()
	} else if keyType == "hash" {
		cmd := service.redisCluster.HGetAll(context.TODO(), key)
		value, err = cmd.Result()
	} else {
		println(keyType)
	}
	valueInfo.Type = keyType
	valueInfo.Value = value
	return
}

func (service *RedisClusterService) Set(key string, value string) (err error) {
	cmd := service.redisCluster.Set(context.TODO(), key, value, time.Duration(0))
	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Sadd(key string, value string) (err error) {
	cmd := service.redisCluster.SAdd(context.TODO(), key, value)
	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Srem(key string, value string) (err error) {
	cmd := service.redisCluster.SRem(context.TODO(), key, value)
	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Lpush(key string, value string) (err error) {
	cmd := service.redisCluster.LPush(context.TODO(), key, value)
	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Rpush(key string, value string) (err error) {
	cmd := service.redisCluster.RPush(context.TODO(), key, value)
	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Lset(key string, index int64, value string) (err error) {
	cmd := service.redisCluster.LSet(context.TODO(), key, index, value)
	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Lrem(key string, count int64, value string) (err error) {
	cmd := service.redisCluster.LRem(context.TODO(), key, count, value)
	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Hset(key string, field string, value string) (err error) {
	cmd := service.redisCluster.HSet(context.TODO(), key, field, value)
	_, err = cmd.Result()
	return
}

func (service *RedisClusterService) Hdel(key string, field string) (err error) {
	cmd := service.redisCluster.HDel(context.TODO(), key, field)
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
	if err != nil {
		return
	}
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
