package redis

import "worker"

func Init() {
	worker_ := &worker.Worker{
		Name:    "redis",
		Text:    "Redis",
		WorkMap: map[string]func(interface{}) (interface{}, error){},
	}
	worker_.WorkMap["get"] = getWork
	worker_.WorkMap["save"] = saveWork
	worker_.WorkMap["keys"] = keysWork
	worker_.WorkMap["delete"] = deleteWork
	worker_.WorkMap["deletePattern"] = deletePatternWork

	worker.AddWorker(worker_)

}
