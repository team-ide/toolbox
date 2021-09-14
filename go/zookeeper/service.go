package zookeeper

import (
	"base"
	"worker"
)

type Automatic struct {
	url               string
	automaticShutdown *worker.AutomaticShutdown
}

func (automatic *Automatic) CreateAutomaticShutdown(automaticShutdown *worker.AutomaticShutdown) error {
	service, err := CreateZKService(automatic.url)
	if err != nil {
		return err
	}
	_, err = service.Exists("/")
	if err != nil {
		service.conn.Close()
		return err
	}
	// 默认10分钟自动关闭
	automaticShutdown.AutomaticShutdown = 10 * 60
	automaticShutdown.Service = service
	automaticShutdown.Stop = func() {
		service.conn.Close()
	}
	automatic.automaticShutdown = automaticShutdown

	return err
}

func getService(url string) (service *ZKService, err error) {
	automatic := &Automatic{
		url: url,
	}
	key := "zookeeper-" + url
	var automaticShutdown *worker.AutomaticShutdown
	automaticShutdown, err = worker.GetAutomaticShutdown(key, automatic.CreateAutomaticShutdown)
	if err != nil {
		return
	}
	automaticShutdown.LastUseTimestamp = base.GetNowTime()
	service = automaticShutdown.Service.(*ZKService)
	return
}
