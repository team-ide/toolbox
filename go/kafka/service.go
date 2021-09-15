package kafka

import (
	"base"
	"worker"
)

type Automatic struct {
	address           string
	automaticShutdown *worker.AutomaticShutdown
}

func (automatic *Automatic) CreateAutomaticShutdown(automaticShutdown *worker.AutomaticShutdown) error {
	service, err := CreateKafkaService(automatic.address)
	if err != nil {
		return err
	}
	_, err = service.GetTopics()
	if err != nil {
		// service.saramaClient.Close()
		return err
	}
	// 默认10分钟自动关闭
	automaticShutdown.AutomaticShutdown = 10 * 60
	automaticShutdown.Service = service
	automaticShutdown.Stop = func() {
		// service.saramaClient.Close()
	}
	automatic.automaticShutdown = automaticShutdown

	return err
}

func getService(address string) (service *KafkaService, err error) {
	automatic := &Automatic{
		address: address,
	}
	key := "kafka-" + address
	var automaticShutdown *worker.AutomaticShutdown
	automaticShutdown, err = worker.GetAutomaticShutdown(key, automatic.CreateAutomaticShutdown)
	if err != nil {
		return
	}
	automaticShutdown.LastUseTimestamp = base.GetNowTime()
	service = automaticShutdown.Service.(*KafkaService)
	return
}
