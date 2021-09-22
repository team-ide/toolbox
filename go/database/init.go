package database_

import "worker"

func Init() {

	worker_ := &worker.Worker{
		Name:    "database",
		Text:    "Database",
		WorkMap: map[string]func(interface{}) (interface{}, error){},
	}
	worker_.WorkMap["databases"] = databasesWork
	worker_.WorkMap["tables"] = tablesWork
	worker_.WorkMap["showCreateDatabase"] = showCreateDatabaseWork
	worker_.WorkMap["showCreateTable"] = showCreateTableWork

	worker.AddWorker(worker_)
}
