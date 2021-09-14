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

type pollRequest struct {
	Address   string `json:"address"`
	GroupId   string `json:"groupId"`
	Topic     string `json:"topic"`
	KeyType   string `json:"keyType"`
	ValueType string `json:"valueType"`
}

type pollResponse struct {
	Msgs []KafkaMessage `json:"msgs"`
}

func pollWork(req interface{}) (res interface{}, err error) {
	request := &pollRequest{}
	response := &pollResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}
	kafkaMsgs, err := service.Poll(request.GroupId, []string{request.Topic})
	if err != nil {
		return nil, err
	}
	msgs := []KafkaMessage{}
	for _, kafkaMsg := range kafkaMsgs {
		msg := KafkaMessage{
			Key:       sarama.StringEncoder(kafkaMsg.Key),
			Value:     sarama.StringEncoder(kafkaMsg.Value),
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
	kafkaMsg := &sarama.ProducerMessage{}
	kafkaMsg.Topic = request.Topic
	kafkaMsg.Key = sarama.StringEncoder(request.Key)
	kafkaMsg.Value = sarama.StringEncoder(request.Value)

	err = service.Push(kafkaMsg)
	if err != nil {
		return nil, err
	}
	res = response
	return
}
