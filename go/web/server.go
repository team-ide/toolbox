package web

import (
	"net/http"
	"worker"
)

func handleOpen(rw http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	workers := []map[string]interface{}{}

	workerCache := worker.WorkerCache
	for _, oneWorker := range workerCache {
		workMap := oneWorker.WorkMap
		workerData := make(map[string]interface{})
		workerData["name"] = oneWorker.Name
		workerData["icon"] = oneWorker.Icon
		workerData["comment"] = oneWorker.Comment
		workerData["text"] = oneWorker.Text

		workerWorks := []map[string]interface{}{}
		for workName, _ := range workMap {
			workData := make(map[string]interface{})

			workData["name"] = workName

			workerWorks = append(workerWorks, workData)
		}
		workerData["works"] = workerWorks
		workers = append(workers, workerData)

	}
	data["workers"] = workers
	outJSON(rw, data, nil)
}

func handleSession(rw http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["A"] = 1
	outJSON(rw, data, nil)
}
