package redis

import (
	"base"
	"fmt"
	"worker"
)

type Automatic struct {
	address           string
	auth              string
	automaticShutdown *worker.AutomaticShutdown
}

func (automatic *Automatic) CreateAutomaticShutdown(automaticShutdown *worker.AutomaticShutdown) error {
	service, err := CreateRedisService(automatic.address, automatic.auth)
	if err != nil {
		return err
	}
	// 默认10分钟自动关闭
	automaticShutdown.AutomaticShutdown = 10 * 60
	automaticShutdown.Service = service
	automaticShutdown.Stop = func() {
		service.pool.Close()
	}
	automatic.automaticShutdown = automaticShutdown

	return err
}

func (automatic *Automatic) CreateClusterAutomaticShutdown(automaticShutdown *worker.AutomaticShutdown) error {
	service, err := CreateRedisClusterService(automatic.address, automatic.auth)
	if err != nil {
		return err
	}
	// 默认10分钟自动关闭
	automaticShutdown.AutomaticShutdown = 10 * 60
	automaticShutdown.Service = service
	automaticShutdown.Stop = func() {
		service.redisCluster.Close()
	}
	automatic.automaticShutdown = automaticShutdown

	return err
}

func getService(address string, auth string, cluster bool) (service Service, err error) {
	automatic := &Automatic{
		address: address,
		auth:    auth,
	}
	if cluster {
		key := "redis-cluster-" + address + "-" + auth + "-" + fmt.Sprint(cluster)
		var automaticShutdown *worker.AutomaticShutdown
		automaticShutdown, err = worker.GetAutomaticShutdown(key, automatic.CreateClusterAutomaticShutdown)
		if err != nil {
			return
		}
		automaticShutdown.LastUseTimestamp = base.GetNowTime()
		service = automaticShutdown.Service.(*RedisClusterService)

	} else {
		key := "redis-" + address + "-" + auth
		var automaticShutdown *worker.AutomaticShutdown
		automaticShutdown, err = worker.GetAutomaticShutdown(key, automatic.CreateAutomaticShutdown)
		if err != nil {
			return
		}
		automaticShutdown.LastUseTimestamp = base.GetNowTime()
		service = automaticShutdown.Service.(*RedisService)
	}
	return
}

type Service interface {
	Keys(pattern string, size int) (count int, keys []string, err error)
	Get(key string) (value string, err error)
	Set(key string, value string) (err error)
	Del(key string) (count int, err error)
	DelPattern(key string) (count int, err error)
}
