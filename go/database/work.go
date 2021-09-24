package database_

import (
	"base"
)

type databasesRequest struct {
	Config DatabaseConfig `json:"config"`
}

type databasesResponse struct {
	Databases []DatabaseInfo `json:"databases"`
}

func databasesWork(req interface{}) (res interface{}, err error) {
	request := &databasesRequest{}
	response := &databasesResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service DatabaseService
	service, err = getService(request.Config)
	if err != nil {
		return
	}
	var databases []DatabaseInfo
	databases, err = service.Databases()
	if err != nil {
		return
	}
	response.Databases = databases
	res = response
	return
}

type tablesRequest struct {
	Config   DatabaseConfig `json:"config"`
	Database string         `json:"database"`
}

type tablesResponse struct {
	Tables []TableInfo `json:"tables"`
}

func tablesWork(req interface{}) (res interface{}, err error) {
	request := &tablesRequest{}
	response := &tablesResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service DatabaseService
	service, err = getService(request.Config)
	if err != nil {
		return
	}
	var tables []TableInfo
	tables, err = service.Tables(request.Database)
	if err != nil {
		return
	}
	response.Tables = tables
	res = response
	return
}

type showCreateDatabaseRequest struct {
	Config   DatabaseConfig `json:"config"`
	Database string         `json:"database"`
}

type showCreateDatabaseResponse struct {
	Create string `json:"create"`
}

func showCreateDatabaseWork(req interface{}) (res interface{}, err error) {
	request := &showCreateDatabaseRequest{}
	response := &showCreateDatabaseResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service DatabaseService
	service, err = getService(request.Config)
	if err != nil {
		return
	}
	var create string
	create, err = service.ShowCreateDatabase(request.Database)
	if err != nil {
		return
	}
	response.Create = create
	res = response
	return
}

type showCreateTableRequest struct {
	Config   DatabaseConfig `json:"config"`
	Database string         `json:"database"`
	Table    string         `json:"table"`
}

type showCreateTableResponse struct {
	Create string `json:"create"`
}

func showCreateTableWork(req interface{}) (res interface{}, err error) {
	request := &showCreateTableRequest{}
	response := &showCreateTableResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service DatabaseService
	service, err = getService(request.Config)
	if err != nil {
		return
	}
	var create string
	create, err = service.ShowCreateTable(request.Database, request.Table)
	if err != nil {
		return
	}
	response.Create = create
	res = response
	return
}

type tableDetailRequest struct {
	Config   DatabaseConfig `json:"config"`
	Database string         `json:"database"`
	Table    string         `json:"table"`
}

type tableDetailResponse struct {
	Table TableDetailInfo `json:"table"`
}

func tableDetailWork(req interface{}) (res interface{}, err error) {
	request := &tableDetailRequest{}
	response := &tableDetailResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service DatabaseService
	service, err = getService(request.Config)
	if err != nil {
		return
	}
	var table TableDetailInfo
	table, err = service.TableDetail(request.Database, request.Table)
	if err != nil {
		return
	}
	response.Table = table
	res = response
	return
}
