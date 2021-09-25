package web

import (
	"base"
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

	enum_map := map[string]interface{}{}

	sqlConditionalOperations := []base.Option{}
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "等于", Value: "="})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不等于", Value: "<>"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "大于", Value: ">"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "大于或等于", Value: ">="})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "小于", Value: "<"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "小于或等于", Value: "<="})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "包含", Value: "like"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不包含", Value: "not like"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "开始以", Value: "like start"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "开始不是以", Value: "not like start"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "结束以", Value: "like end"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "结束不是以", Value: "not like end"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "是null", Value: "is null"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不是null", Value: "is not null"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "是空", Value: "is empty"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不是空", Value: "is not empty"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "介于", Value: "between"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不介于", Value: "not between"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "在列表", Value: "in"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不在列表", Value: "not in"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "自定义", Value: "custom"})

	enum_map["sqlConditionalOperations"] = sqlConditionalOperations

	data["enum_map"] = enum_map

	outJSON(rw, data, nil)
}

func handleSession(rw http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["A"] = 1
	outJSON(rw, data, nil)
}
