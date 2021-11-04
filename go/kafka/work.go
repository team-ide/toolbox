package kafka

import (
	"base"
	"encoding/binary"
	"strconv"

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
	Key       interface{} `json:"key"`
	Value     interface{} `json:"value"`
	Topic     string      `json:"topic"`
	Partition int32       `json:"partition"`
	Offset    int64       `json:"offset"`
	Headers   []Header    `json:"headers"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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
		} else if request.KeyType == "Long" {
			if len(kafkaMsg.Key) == 8 {
				key = uint64(binary.BigEndian.Uint64(kafkaMsg.Key))
			} else {
				key = sarama.StringEncoder(kafkaMsg.Key)
			}
		} else {
			key = sarama.ByteEncoder(kafkaMsg.Key)
		}
		if request.ValueType == "String" {
			value = sarama.StringEncoder(kafkaMsg.Value)
		} else if request.ValueType == "Long" {
			if len(kafkaMsg.Value) == 8 {
				value = uint64(binary.BigEndian.Uint64(kafkaMsg.Value))
			} else {
				value = sarama.StringEncoder(kafkaMsg.Value)
			}
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
				msg.Headers = append(msg.Headers, Header{Key: string(header.Key), Value: string(header.Value)})
			}
		}
		msgs = append(msgs, msg)
	}
	response.Msgs = msgs
	res = response
	return
}

type pushRequest struct {
	Address   string   `json:"address"`
	Topic     string   `json:"topic"`
	KeyType   string   `json:"keyType"`
	ValueType string   `json:"valueType"`
	Headers   []Header `json:"headers"`
	Key       string   `json:"key"`
	Value     string   `json:"value"`
	Partition int32    `json:"partition"`
	Offset    int64    `json:"offset"`
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
	if request.Key != "" {
		if request.KeyType == "String" {
			key = sarama.StringEncoder(request.Key)
		} else if request.KeyType == "Long" {
			longV, err := strconv.ParseInt(request.Key, 10, 64)
			if err != nil {
				return nil, err
			}
			uintV := uint64(longV)
			bytes := make([]byte, 8)
			binary.BigEndian.PutUint64(bytes, uintV)
			key = sarama.ByteEncoder(bytes)
		} else {
			key = sarama.ByteEncoder(request.Key)
		}
	}
	if request.Value != "" {
		if request.ValueType == "String" {
			value = sarama.StringEncoder(request.Value)
		} else if request.ValueType == "Long" {
			longV, err := strconv.ParseInt(request.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			uintV := uint64(longV)
			bytes := make([]byte, 8)
			binary.BigEndian.PutUint64(bytes, uintV)
			value = sarama.ByteEncoder(bytes)
		} else {
			value = sarama.ByteEncoder(request.Value)
		}
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
	if request.Headers != nil {
		for _, one := range request.Headers {
			kafkaMsg.Headers = append(kafkaMsg.Headers, sarama.RecordHeader{
				Key:   []byte(one.Key),
				Value: []byte(one.Value),
			})
		}
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

	err = service.MarkOffset(request.GroupId, request.Topic, request.Partition, request.Offset)
	if err != nil {
		return nil, err
	}
	res = response
	return
}

type resetRequest struct {
	Address   string `json:"address"`
	GroupId   string `json:"groupId"`
	Topic     string `json:"topic"`
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
}

type resetResponse struct {
}

func resetWork(req interface{}) (res interface{}, err error) {
	request := &resetRequest{}
	response := &resetResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}

	err = service.ResetOffset(request.GroupId, request.Topic, request.Partition, request.Offset)
	if err != nil {
		return nil, err
	}
	res = response
	return
}

type deleteTopicRequest struct {
	Address string `json:"address"`
	Topic   string `json:"topic"`
}

type deleteTopicResponse struct {
}

func deleteTopicWork(req interface{}) (res interface{}, err error) {
	request := &deleteTopicRequest{}
	response := &deleteTopicResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}

	err = service.DeleteTopic(request.Topic)
	if err != nil {
		return nil, err
	}
	res = response
	return
}

type createPartitionsRequest struct {
	Address string `json:"address"`
	Topic   string `json:"topic"`
	Count   int32  `json:"count"`
}

type createPartitionsResponse struct {
}

func createPartitionsWork(req interface{}) (res interface{}, err error) {
	request := &createPartitionsRequest{}
	response := &createPartitionsResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}

	err = service.CreatePartitions(request.Topic, request.Count)
	if err != nil {
		return nil, err
	}
	res = response
	return
}

type deleteRecordsRequest struct {
	Address   string `json:"address"`
	Topic     string `json:"topic"`
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
}

type deleteRecordsResponse struct {
}

func deleteRecordsWork(req interface{}) (res interface{}, err error) {
	request := &deleteRecordsRequest{}
	response := &deleteRecordsResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service *KafkaService
	service, err = getService(request.Address)
	if err != nil {
		return
	}

	partitionOffsets := make(map[int32]int64)
	partitionOffsets[request.Partition] = request.Offset
	err = service.DeleteRecords(request.Topic, partitionOffsets)
	if err != nil {
		return nil, err
	}
	res = response
	return
}
