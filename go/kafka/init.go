package kafka

import "worker"

func Init() {

	worker_ := &worker.Worker{
		Name:    "kafka",
		Text:    "Kafka",
		WorkMap: map[string]func(interface{}) (interface{}, error){},
	}
	worker_.WorkMap["topics"] = topicsWork
	worker_.WorkMap["pull"] = pullWork
	worker_.WorkMap["push"] = pushWork
	worker_.WorkMap["commit"] = commitWork
	worker_.WorkMap["reset"] = resetWork
	worker_.WorkMap["deleteTopic"] = deleteTopicWork
	worker_.WorkMap["createPartitions"] = createPartitionsWork
	worker_.WorkMap["deleteRecords"] = deleteRecordsWork

	worker.AddWorker(worker_)
}
