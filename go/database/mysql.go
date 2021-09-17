package database_

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {
}

type MysqlService struct {
	config DatabaseConfig
	db     *sqlx.DB
}

func CreateMysqlService(config DatabaseConfig) (service *MysqlService, err error) {
	service = &MysqlService{
		config: config,
	}
	err = service.init()
	return
}

func (service *MysqlService) init() (err error) {
	url := fmt.Sprint(service.config.Username, ":", service.config.Password, "@tcp(", service.config.Host, ":", service.config.Port, ")/?charset=utf8")
	var db *sqlx.DB
	db, err = sqlx.Open("mysql", url)
	service.db = db
	return
}

func (service *MysqlService) Open() (err error) {
	err = service.db.Ping()
	return
}

func (service *MysqlService) Close() (err error) {
	err = service.db.Close()
	return
}

func (service *MysqlService) Databases() (databases []DatabaseInfo, err error) {
	sqlParam := SqlParam{
		Sql:    "show databases",
		Params: []interface{}{},
	}
	res, err := service.Query(sqlParam)
	if err != nil {
		return
	}
	for _, one := range res {
		info := DatabaseInfo{}
		info.Name = string(one["Database"])
		databases = append(databases, info)
	}
	return
}

func (service *MysqlService) Tables(database string) (tables []TableInfo, err error) {

	sql_ := `select TABLE_NAME name,TABLE_COMMENT comment from information_schema.tables  where TABLE_SCHEMA=? and TABLE_TYPE='BASE TABLE'`
	sqlParam := SqlParam{
		Sql:    sql_,
		Params: []interface{}{database},
	}
	res, err := service.Query(sqlParam)
	if err != nil {
		return
	}
	for _, one := range res {
		info := TableInfo{
			Name:    string(one["name"]),
			Comment: string(one["comment"]),
		}
		tables = append(tables, info)
	}
	return
}

func (service *MysqlService) Query(sqlParam SqlParam) (res []map[string][]byte, err error) {
	rows, err := service.db.Query(sqlParam.Sql, sqlParam.Params...)
	if err != nil {
		return
	}
	res, err = ResultToMap(rows)
	if err != nil {
		return
	}
	rows.Close()
	return
}

func (service *MysqlService) Insert(sqlParam SqlParam) (rowsAffected int64, err error) {

	result, err := service.db.Exec(sqlParam.Sql, sqlParam.Params...)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}

func (service *MysqlService) Update(sqlParam SqlParam) (rowsAffected int64, err error) {
	result, err := service.db.Exec(sqlParam.Sql, sqlParam.Params...)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}

func (service *MysqlService) Delete(sqlParam SqlParam) (rowsAffected int64, err error) {
	result, err := service.db.Exec(sqlParam.Sql, sqlParam.Params...)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}
