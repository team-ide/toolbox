package kafka

import (
	"base"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func CreateKafkaService(address string) (service *KafkaService, err error) {
	service = &KafkaService{
		address: address,
	}
	return
}

type KafkaService struct {
	address string
}

func (service *KafkaService) getClient() (saramaClient sarama.Client, err error) {
	SaramaConfig := sarama.NewConfig()
	SaramaConfig.Consumer.Return.Errors = true
	SaramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	SaramaConfig.Consumer.MaxWaitTime = time.Second * 1
	addrs := strings.Split(service.address, ",")
	saramaClient, err = sarama.NewClient(addrs, SaramaConfig)
	if err != nil {
		return
	}
	return
}

func (service *KafkaService) GetTopics() (topics []string, err error) {
	var saramaClient sarama.Client
	saramaClient, err = service.getClient()
	if err != nil {
		return
	}
	defer saramaClient.Close()
	topics, err = saramaClient.Topics()
	return
}

func (service *KafkaService) Pull(groupId string, topics []string) (msgs []*sarama.ConsumerMessage, err error) {
	var saramaClient sarama.Client
	saramaClient, err = service.getClient()
	if err != nil {
		return
	}
	defer saramaClient.Close()
	group, err := sarama.NewConsumerGroupFromClient(groupId, saramaClient)
	if err != nil {
		return
	}
	handler := &consumerGroupHandler{}
	go func() {
		ctx := context.Background()
		err = group.Consume(ctx, topics, handler)
		if err != nil {
			fmt.Println("group.Consume error:", err)
		}
	}()
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
	// go func() {
	err = group.Close()
	if err != nil {
		fmt.Println("group.Close error:", err)
		return
	}
	// }()
	msgs = handler.msgs
	return
}

type consumerGroupHandler struct {
	msgs []*sarama.ConsumerMessage
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (handler *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	chanMessages := claim.Messages()
	for msg := range chanMessages {
		handler.msgs = append(handler.msgs, msg)
	}
	return nil
}

func (service *KafkaService) MarkOffset(groupId string, topic string, partition int32, offset int64) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = service.getClient()
	if err != nil {
		return
	}
	defer saramaClient.Close()
	offsetManager, err := sarama.NewOffsetManagerFromClient(groupId, saramaClient)
	if err != nil {
		return
	}
	partitionOffsetManager, err := offsetManager.ManagePartition(topic, partition)
	if err != nil {
		return
	}
	partitionOffsetManager.MarkOffset(offset, "")
	err = offsetManager.Close()
	return
}

func (service *KafkaService) ResetOffset(groupId string, topic string, partition int32, offset int64) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = service.getClient()
	if err != nil {
		return
	}
	defer saramaClient.Close()
	offsetManager, err := sarama.NewOffsetManagerFromClient(groupId, saramaClient)
	if err != nil {
		return
	}
	partitionOffsetManager, err := offsetManager.ManagePartition(topic, partition)
	if err != nil {
		return
	}
	defer offsetManager.Close()
	partitionOffsetManager.ResetOffset(offset, "")
	return
}

func (service *KafkaService) CreatePartitions(topic string, count int32) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = service.getClient()
	if err != nil {
		return
	}
	defer saramaClient.Close()
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer admin.Close()

	err = admin.CreatePartitions(topic, count, nil, false)

	return
}

func (service *KafkaService) DeleteTopic(topic string) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = service.getClient()
	if err != nil {
		return
	}
	defer saramaClient.Close()
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer admin.Close()

	err = admin.DeleteTopic(topic)

	return
}

func (service *KafkaService) DeleteConsumerGroup(groupId string) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = service.getClient()
	if err != nil {
		return
	}
	defer saramaClient.Close()
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer admin.Close()

	err = admin.DeleteConsumerGroup(groupId)

	return
}

func (service *KafkaService) DeleteRecords(topic string, partitionOffsets map[int32]int64) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = service.getClient()
	if err != nil {
		return
	}
	defer saramaClient.Close()
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer admin.Close()

	err = admin.DeleteRecords(topic, partitionOffsets)

	return
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
