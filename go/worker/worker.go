package worker

import (
	"base"
	"sync"
	"time"
)

type Worker struct {
	Name    string
	Text    string
	Icon    string
	Comment string
	WorkMap map[string]func(interface{}) (interface{}, error)
}

var (
	WorkerCache                   map[string]*Worker
	toolboxAutomaticShutdownCache map[string]*ToolboxAutomaticShutdown
	lock                          sync.Mutex
)

func init() {
	WorkerCache = map[string]*Worker{}
	toolboxAutomaticShutdownCache = map[string]*ToolboxAutomaticShutdown{}
	go startToolboxAutomaticShutdownTimer()
}

func AddWorker(worker *Worker) {
	name := worker.Name
	WorkerCache[name] = worker
}

func GetToolboxAutomaticShutdown(key string, create func() *ToolboxAutomaticShutdown) *ToolboxAutomaticShutdown {
	lock.Lock()
	defer lock.Unlock()
	res, ok := toolboxAutomaticShutdownCache[key]
	if !ok {
		res = create()
		toolboxAutomaticShutdownCache[key] = res
	}
	return res
}

type ToolboxAutomaticShutdown interface {
	AutomaticShutdown() int64
	LastUseTimestamp() int64
	stop()
	started()
}

func startToolboxAutomaticShutdownTimer() {
	time.Sleep(1 * time.Second)
	cache := toolboxAutomaticShutdownCache
	nowTime := base.GetNowTime()
	for _, one := range cache {
		one_ := *one
		if one_.AutomaticShutdown() <= 0 {
			continue
		}
		lastTime := nowTime - one_.LastUseTimestamp()
		if lastTime >= one_.AutomaticShutdown() {
			one_.stop()
		}

	}
}
