package web

import (
	"config"
	"fmt"

	"net/http"
	"strings"
	"worker"

	"github.com/gorilla/mux"
)

type WorkerWork struct {
	worker worker.Worker
	work   worker.Work
}

func (workerWork *WorkerWork) handle(rw http.ResponseWriter, r *http.Request) {
	body, err := getPostBody(r)
	if err != nil {
		outJSON(rw, nil, err)
		return
	}
	res, err := workerWork.work.Work(body)
	outJSON(rw, res, err)
}

func StartServer() {

	r := mux.NewRouter()

	workerCache := worker.WorkerCache
	for workerName, oneWorker := range workerCache {
		workMap := oneWorker.WorkMap()
		for workName, oneWork := range workMap {
			workerWork := WorkerWork{
				worker: oneWorker,
				work:   oneWork,
			}
			requestMapping := fmt.Sprint("/server/", workerName, "/", workName)
			r.HandleFunc(requestMapping, workerWork.handle)
		}
	}

	//静态请求，由AssetFS统一处理。
	for k := range _bintree.Children {
		name := k
		path := "/" + name
		if name == "index.html" {
			path = "/"
		}
		r.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(name, ".html") {
				rw.Header().Set("Content-Type", "text/html")
			} else if strings.HasSuffix(name, ".css") {
				rw.Header().Set("Content-Type", "text/css")
			} else if strings.HasSuffix(name, ".js") {
				rw.Header().Set("Content-Type", "application/javascript")
			}
			bytes, _ := Asset(name)
			_, _ = rw.Write(bytes)
		})
	}

	host := config.Config.Server.Host
	port := config.Config.Server.Port

	httpServer := fmt.Sprint(host, ":", port)

	err := http.ListenAndServe(httpServer, r)
	if err != nil {
		panic(err)
	}
}
