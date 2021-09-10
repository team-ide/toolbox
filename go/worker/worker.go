package worker

type Worker interface {
	Name() string
	Text() string
	Icon() string
	Comment() string
	WorkMap() map[string]Work
}

type Work interface {
	Work(interface{}) (interface{}, error)
}

var (
	WorkerCache map[string]Worker
)

func init() {
	WorkerCache = map[string]Worker{}
}

func AddWorker(worker Worker) {
	name := worker.Name()
	WorkerCache[name] = worker
}
