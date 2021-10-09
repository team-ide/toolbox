package redis

import (
	"config"
	"worker"
)

func Init() {
	worker_ := &worker.Worker{
		Name:    "redis",
		Text:    "Redis",
		Configs: config.Config.Redis,
		WorkMap: map[string]func(interface{}) (interface{}, error){},
	}
	worker_.WorkMap["get"] = getWork
	worker_.WorkMap["do"] = doWork
	worker_.WorkMap["keys"] = keysWork
	worker_.WorkMap["delete"] = deleteWork
	worker_.WorkMap["deletePattern"] = deletePatternWork

	worker.AddWorker(worker_)

}
