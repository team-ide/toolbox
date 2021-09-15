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

	worker.AddWorker(worker_)
}
