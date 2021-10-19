package database_

import (
	"fmt"
	"strconv"
	"strings"

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
		keys := make([]string, 0, len(one))
		for k := range one {
			keys = append(keys, k)
		}
		info := DatabaseInfo{}
		info.Name = string(one[keys[0]])
		databases = append(databases, info)
	}
	return
}

func (service *MysqlService) ShowCreateDatabase(database string) (create string, err error) {
	sqlParam := SqlParam{
		Sql:    "show create database `" + database + "`",
		Params: []interface{}{},
	}
	res, err := service.Query(sqlParam)
	if err != nil {
		return
	}
	for _, one := range res {
		keys := make([]string, 0, len(one))
		for k := range one {
			keys = append(keys, k)
		}
		create = string(one[keys[1]])
	}
	return
}

func (service *MysqlService) ShowCreateTable(database string, table string) (create string, err error) {
	sqlParam := SqlParam{
		Sql:    "show create table `" + database + "`.`" + table + "`",
		Params: []interface{}{},
	}
	res, err := service.Query(sqlParam)
	if err != nil {
		return
	}
	for _, one := range res {
		keys := make([]string, 0, len(one))
		for k := range one {
			keys = append(keys, k)
		}

		create = string(one[keys[1]])
	}
	return
}

func (service *MysqlService) Tables(database string) (tables []TableInfo, err error) {

	sql_ := "show table status from `" + database + "`"
	sqlParam := SqlParam{
		Sql:    sql_,
		Params: []interface{}{},
	}
	res, err := service.Query(sqlParam)
	if err != nil {
		return
	}
	for _, one := range res {
		info := TableInfo{
			Name:    string(one["Name"]),
			Comment: string(one["Comment"]),
		}
		tables = append(tables, info)
	}
	return
}

func (service *MysqlService) TableDetail(database string, table string) (tableDetail TableDetailInfo, err error) {

	sql_ := "show table status from `" + database + "` where Name=?"
	sqlParam := SqlParam{
		Sql:    sql_,
		Params: []interface{}{table},
	}
	res, err := service.Query(sqlParam)
	if err != nil {
		return
	}
	if len(res) == 0 {
		return
	}
	tableDetail = TableDetailInfo{
		Name:    string(res[0]["Name"]),
		Comment: string(res[0]["Comment"]),
	}
	var columns []TableColumnInfo
	columns, err = service.TableColumns(database, table)
	if err != nil {
		return
	}
	tableDetail.Columns = columns

	var indexs []TableIndexInfo
	indexs, err = service.TableIndexs(database, table)
	if err != nil {
		return
	}
	tableDetail.Indexs = indexs
	return
}

func (service *MysqlService) TableColumns(database string, table string) (columns []TableColumnInfo, err error) {

	sql_ := "show full columns from  `" + database + "`.`" + table + "`"
	sqlParam := SqlParam{
		Sql:    sql_,
		Params: []interface{}{},
	}
	res, err := service.Query(sqlParam)
	if err != nil {
		return
	}
	for _, one := range res {
		info := TableColumnInfo{
			Name:    string(one["Field"]),
			Comment: string(one["Comment"]),
		}
		columnTypeStr := string(one["Type"])
		columnType := columnTypeStr

		if strings.Contains(columnTypeStr, "(") {
			columnType = columnTypeStr[0:strings.Index(columnTypeStr, "(")]
			lengthStr := columnTypeStr[strings.Index(columnTypeStr, "(")+1:]
			if strings.Contains(lengthStr, ",") {
				length, _ := strconv.Atoi(lengthStr[0:strings.Index(lengthStr, ",")])
				decimal, _ := strconv.Atoi(lengthStr[strings.Index(lengthStr, ",")+1:])
				info.Length = length
				info.Decimal = decimal
			} else {
				length, _ := strconv.Atoi(lengthStr)
				info.Length = length
			}
		}
		info.Type = columnType
		columns = append(columns, info)
	}
	return
}

func (service *MysqlService) TableIndexs(database string, table string) (indexs []TableIndexInfo, err error) {

	sql_ := "show indexes from  `" + database + "`.`" + table + "`"
	sqlParam := SqlParam{
		Sql:    sql_,
		Params: []interface{}{},
	}
	res, err := service.Query(sqlParam)
	if err != nil {
		return
	}
	for _, one := range res {
		Key_name := string(one["Key_name"])
		if Key_name == "PRIMARY" {
			continue
		}
		info := TableIndexInfo{
			Name:    Key_name,
			Comment: string(one["Comment"]),
		}
		indexs = append(indexs, info)
	}
	return
}

func (service *MysqlService) Datas(datasParam DatasParam) (datasResult DatasResult, err error) {

	sql_ := "SELECT "
	countSql := "SELECT "
	params_ := []interface{}{}

	columnMap := make(map[string]TableColumnInfo)

	for _, column := range datasParam.Columns {
		columnMap[column.Name] = column
		sql_ += "`" + column.Name + "`,"
	}
	sql_ = sql_[0 : len(sql_)-1]

	sql_ += " FROM `" + datasParam.Database + "`.`" + datasParam.Table + "` "
	countSql += " COUNT(1) AS total FROM `" + datasParam.Database + "`.`" + datasParam.Table + "` "

	if len(datasParam.Wheres) > 0 {
		whereSql := "WHERE "
		for index, where := range datasParam.Wheres {
			value := where.Value
			whereSql += "`" + where.Name + "` "
			switch where.SqlConditionalOperation {
			case "like":
				whereSql += "LIKE '%" + value + "%' "
			case "not like":
				whereSql += "NOT LIKE '%" + value + "%' "
			case "like start":
				whereSql += "LIKE '" + value + "%' "
			case "not like start":
				whereSql += "NOT LIKE '" + value + "%' "
			case "like end":
				whereSql += "LIKE '%" + value + "' "
			case "not like end":
				whereSql += "NOT LIKE '%" + value + "' "
			case "is null":
				whereSql += "IS NULL "
			case "is not null":
				whereSql += "IS NOT NULL "
			case "is empty":
				whereSql += "= '' "
			case "is not empty":
				whereSql += "<> '' "
			case "between":
				whereSql += "BETWEEN " + "'" + where.Before + "' AND '" + where.After + "' "
			case "not between":
				whereSql += "NOT BETWEEN " + "'" + where.Before + "' AND '" + where.After + "' "
			case "in":
				whereSql += "IN " + "(" + value + ") "
			case "not in":
				whereSql += "NOT IN " + "(" + value + ") "
			default:
				whereSql += where.SqlConditionalOperation + " '" + value + "' "
			}
			// params_ = append(params_, where.Value)
			if index < len(datasParam.Wheres)-1 {
				whereSql += " " + where.AndOr + " "
			}
		}
		sql_ += whereSql
		countSql += whereSql
	}
	sql_ = fmt.Sprint(sql_, " LIMIT ", (datasParam.PageIndex-1)*datasParam.PageSize, ",", datasParam.PageSize)

	totalRes, err := service.Query(SqlParam{
		Sql:    countSql,
		Params: params_,
	})
	if err != nil {
		return
	}
	total := string(totalRes[0]["total"])
	res, err := service.Query(SqlParam{
		Sql:    sql_,
		Params: params_,
	})
	if err != nil {
		return
	}
	datas := []map[string]interface{}{}
	for _, one := range res {
		data := make(map[string]interface{})
		for key, value := range one {
			column := columnMap[key]
			var value_ interface{}
			if value != nil {
				if column.Type != "" {
					value_ = string(value)
				} else {
					value_ = value
				}
			}
			data[key] = value_
		}
		datas = append(datas, data)
	}
	datasResult.Sql = sql_
	datasResult.Total = total
	datasResult.Params = params_
	datasResult.Datas = datas
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
