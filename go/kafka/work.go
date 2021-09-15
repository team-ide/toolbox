package kafka

import (
	"base"

	"github.com/Shopify/sarama"
)

type topicsRequest struct {
	Address string `json:"address"`
}

type topicsResponse struct {
	Topics []string `json:"topics"`
}

func topicsWork(req interface{}) (res interface{}, err error) {
	request := &topicsRequest{}
	response := &topicsResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}
	var topics []string
	topics, err = service.GetTopics()
	if err != nil {
		return
	}
	response.Topics = topics
	res = response
	return
}

type KafkaMessage struct {
	Key       interface{}       `json:"key"`
	Value     interface{}       `json:"value"`
	Topic     string            `json:"topic"`
	Partition int32             `json:"partition"`
	Offset    int64             `json:"offset"`
	Header    map[string]string `json:"header"`
}

type pullRequest struct {
	Address   string `json:"address"`
	GroupId   string `json:"groupId"`
	Topic     string `json:"topic"`
	KeyType   string `json:"keyType"`
	ValueType string `json:"valueType"`
}

type pullResponse struct {
	Msgs []KafkaMessage `json:"msgs"`
}

func pullWork(req interface{}) (res interface{}, err error) {
	request := &pullRequest{}
	response := &pullResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}
	kafkaMsgs, err := service.Pull(request.GroupId, []string{request.Topic})
	if err != nil {
		return nil, err
	}
	msgs := []KafkaMessage{}
	for _, kafkaMsg := range kafkaMsgs {
		var key interface{}
		var value interface{}
		if request.KeyType == "String" {
			key = sarama.StringEncoder(kafkaMsg.Key)
		} else {
			key = sarama.ByteEncoder(kafkaMsg.Key)
		}
		if request.ValueType == "String" {
			value = sarama.StringEncoder(kafkaMsg.Value)
		} else {
			value = sarama.ByteEncoder(kafkaMsg.Value)
		}
		msg := KafkaMessage{
			Key:       key,
			Value:     value,
			Topic:     kafkaMsg.Topic,
			Partition: kafkaMsg.Partition,
			Offset:    kafkaMsg.Offset,
		}
		if kafkaMsg.Headers != nil {
			for _, header := range kafkaMsg.Headers {
				msg.Header[string(header.Key)] = string(header.Value)
			}
		}
		msgs = append(msgs, msg)
	}
	response.Msgs = msgs
	res = response
	return
}

type pushRequest struct {
	Address   string `json:"address"`
	Topic     string `json:"topic"`
	KeyType   string `json:"keyType"`
	ValueType string `json:"valueType"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
}

type pushResponse struct {
}

func pushWork(req interface{}) (res interface{}, err error) {
	request := &pushRequest{}
	response := &pushResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}

	var key sarama.Encoder
	var value sarama.Encoder
	if request.KeyType == "String" {
		key = sarama.StringEncoder(request.Key)
	} else {
		key = sarama.ByteEncoder(request.Key)
	}
	if request.ValueType == "String" {
		value = sarama.StringEncoder(request.Value)
	} else {
		value = sarama.ByteEncoder(request.Value)
	}

	kafkaMsg := &sarama.ProducerMessage{}
	kafkaMsg.Topic = request.Topic
	kafkaMsg.Key = key
	kafkaMsg.Value = value
	if request.Partition >= 0 {
		kafkaMsg.Partition = request.Partition
	}
	if request.Offset >= 0 {
		kafkaMsg.Offset = request.Offset
	}

	err = service.Push(kafkaMsg)
	if err != nil {
		return nil, err
	}
	res = response
	return
}

type commitRequest struct {
	Address   string `json:"address"`
	GroupId   string `json:"groupId"`
	Topic     string `json:"topic"`
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
}

type commitResponse struct {
}

func commitWork(req interface{}) (res interface{}, err error) {
	request := &commitRequest{}
	response := &commitResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}

	err = service.Commit(request.GroupId, request.Topic, request.Partition, request.Offset)
	if err != nil {
		return nil, err
	}
	res = response
	return
}
