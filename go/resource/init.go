package resource

import (
	"worker"
)

func Init() {
	worker_ := &worker.Worker{
		Name:    "resource",
		Text:    "Resource",
		WorkMap: map[string]func(interface{}) (interface{}, error){},
	}
	worker_.WorkMap["list"] = listWork

	worker.AddWorker(worker_)

}
