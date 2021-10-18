package database_

import (
	"base"
)

type checkConnectRequest struct {
	Config DatabaseConfig `json:"config"`
}

type checkConnectResponse struct {
}

func checkConnectWork(req interface{}) (res interface{}, err error) {
	request := &checkConnectRequest{}
	response := &checkConnectResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service DatabaseService
	service, err = getService(request.Config)
	if err != nil {
		return
	}
	res = service.Open()
	res = response
	return
}

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

type datasRequest struct {
	Config    DatabaseConfig    `json:"config"`
	Database  string            `json:"database"`
	Table     string            `json:"table"`
	Columns   []TableColumnInfo `json:"columns"`
	Wheres    []Where           `json:"wheres"`
	PageIndex int               `json:"pageIndex"`
	PageSize  int               `json:"pageSize"`
}

type datasResponse struct {
	Sql    string                   `json:"sql"`
	Params []interface{}            `json:"params"`
	Datas  []map[string]interface{} `json:"datas"`
}

func datasWork(req interface{}) (res interface{}, err error) {
	request := &datasRequest{}
	response := &datasResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service DatabaseService
	service, err = getService(request.Config)
	if err != nil {
		return
	}
	var datasRequest DatasResult
	datasRequest, err = service.Datas(DatasParam{
		Database:  request.Database,
		Table:     request.Table,
		Columns:   request.Columns,
		Wheres:    request.Wheres,
		PageIndex: request.PageIndex,
		PageSize:  request.PageSize,
	})
	if err != nil {
		return
	}
	response.Sql = datasRequest.Sql
	response.Params = datasRequest.Params
	response.Datas = datasRequest.Datas
	res = response
	return
}
