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
	worker *worker.Worker
	work   func(interface{}) (interface{}, error)
}

func (workerWork *WorkerWork) handle(rw http.ResponseWriter, r *http.Request) {
	body, err := getPostBody(r)
	if err != nil {
		outJSON(rw, nil, err)
		return
	}

	res, err := workerWork.work(body)
	outJSON(rw, res, err)
}

func handleIndex(rw http.ResponseWriter, r *http.Request) {
	context := config.Config.Server.Context
	if context == "" || context == "/" {
		context = ""
	}
	rw.Header().Set("refresh", "0.1;"+context+"/")
	rw.WriteHeader(200)
}

func StartServer() {

	context := config.Config.Server.Context
	if context == "" || context == "/" {
		context = ""
	}

	r := mux.NewRouter()

	r.HandleFunc(context+"", handleIndex)
	r.HandleFunc(context+"/server/open", handleOpen)
	r.HandleFunc(context+"/server/session", handleSession)
	r.HandleFunc(context+"/server/login", handleLogin)
	r.HandleFunc(context+"/server/logout", handleLogout)

	workerCache := worker.WorkerCache
	for workerName, oneWorker := range workerCache {
		workMap := oneWorker.WorkMap
		for workName, oneWork := range workMap {
			workerWork := WorkerWork{
				worker: oneWorker,
				work:   oneWork,
			}
			requestMapping := fmt.Sprint(context, "/server/worker/", workerName, "/", workName)
			r.HandleFunc(requestMapping, workerWork.handle)
		}
	}

	//静态请求，由AssetFS统一处理。
	for k := range _bintree.Children {
		name := k
		path := context + "/" + name
		if name == "index.html" {
			path = context + "/"
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
	if context != "/" {
		context += "/"
	}
	println("url:http://" + httpServer + context)
	err := http.ListenAndServe(httpServer, r)
	if err != nil {
		panic(err)
	}
}
