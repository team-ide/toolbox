package kafka

import (
	"base"
	"context"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func CreateKafkaService(address string) (service *KafkaService, err error) {
	service = &KafkaService{
		address: address,
	}
	err = service.init()
	return
}

type KafkaService struct {
	address      string
	saramaClient sarama.Client
}

func (service *KafkaService) init() (err error) {
	SaramaConfig := sarama.NewConfig()
	SaramaConfig.Consumer.Return.Errors = true
	SaramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	SaramaConfig.Consumer.MaxWaitTime = time.Second * 5
	addrs := strings.Split(service.address, ",")
	var saramaClient sarama.Client
	saramaClient, err = sarama.NewClient(addrs, SaramaConfig)
	if err != nil {
		return
	}
	service.saramaClient = saramaClient
	return
}

func (service *KafkaService) GetTopics() (topics []string, err error) {
	topics, err = service.saramaClient.Topics()
	return
}

func (service *KafkaService) Poll(groupId string, topics []string) (msgs []*sarama.ConsumerMessage, err error) {
	group, err := sarama.NewConsumerGroupFromClient(groupId, service.saramaClient)
	if err != nil {
		return
	}
	handler := &consumerGroupHandler{}
	go func(handler *consumerGroupHandler) {
		ctx := context.TODO()
		err = group.Consume(ctx, topics, handler)
	}(handler)
	startTime := base.GetNowTime()
	for {
		time.Sleep(200 * time.Millisecond)
		if len(handler.msgs) > 0 {
			break
		}
		nowTime := base.GetNowTime()
		if nowTime-startTime > 5*1000 {
			break
		}
	}
	go func(group sarama.ConsumerGroup) {
		err = group.Close()
		if err != nil {
			return
		}
	}(group)
	msgs = handler.msgs
	return
}

type consumerGroupHandler struct {
	msgs []*sarama.ConsumerMessage
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (handler *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		handler.msgs = append(handler.msgs, msg)
	}
	return nil
}

//创建生产者
func (service *KafkaService) NewSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 3
	var err error
	syncProducer, err := sarama.NewSyncProducer(strings.Split(service.address, ","), config)
	if err != nil {
		return nil, err
	}
	return syncProducer, nil
}

//推送消息到kafka
func (service *KafkaService) Push(msg *sarama.ProducerMessage) (err error) {
	syncProducer, err := service.NewSyncProducer()
	if err != nil {
		return
	}
	defer syncProducer.Close()
	msg.Timestamp = time.Now()
	_, _, err = syncProducer.SendMessage(msg)
	return err
}
