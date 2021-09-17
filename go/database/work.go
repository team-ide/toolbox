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
