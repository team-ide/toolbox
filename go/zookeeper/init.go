package zookeeper

import (
	"worker"
)

func Init() {
	worker_ := &worker.Worker{
		Name:    "zookeeper",
		Text:    "Zookeeper",
		WorkMap: map[string]func(interface{}) (interface{}, error){},
	}
	worker_.WorkMap["get"] = getWork
	worker_.WorkMap["save"] = saveWork
	worker_.WorkMap["getChildren"] = getChildrenWork
	worker_.WorkMap["delete"] = deleteWork

	worker.AddWorker(worker_)
}
