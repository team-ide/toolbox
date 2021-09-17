package database_

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type SqlParam struct {
	Sql    string        `json:"sql,omitempty"`
	Params []interface{} `json:"params,omitempty"`
}

func GetFieldTypeValue(t reflect.Type, v reflect.Value) interface{} {
	var value interface{}
	switch t.Name() {
	case "string":
		value = v.String()
	case "bool":
		value = v.Bool()
	default:
		value = v.Int()
	}
	return value
}

func InsertSqlByBean(bean interface{}, table string) (sqlParam SqlParam) {
	params := []interface{}{}
	refValue := reflect.ValueOf(bean) // value
	refType := reflect.TypeOf(bean)   // type
	fieldCount := refType.NumField()  // field count
	insertColumns := ""
	insertValues := ""
	for i := 0; i < fieldCount; i++ {
		fieldType := refType.Field(i)         // field type
		fieldValue := refValue.Field(i)       // field vlaue
		column := fieldType.Tag.Get("column") // field tag
		if column != "" {
			insertColumns += column + ","
			value := GetFieldTypeValue(fieldType.Type, fieldValue)
			switch fieldType.Type.Name() {
			case "string":
				if value == "" {
					insertValues += "NULL,"
					continue
				}
				insertValues += "?,"
				params = append(params, value)
			default:
				insertValues += "?,"
				params = append(params, value)
			}
		}
	}
	insertColumns = insertColumns[0 : len(insertColumns)-1]
	insertValues = insertValues[0 : len(insertValues)-1]
	sql := " INSERT INTO " + table + "(" + insertColumns + ") VALUES (" + insertValues + ") "
	sqlParam = SqlParam{sql, params}
	return sqlParam
}

func GetColumnSqlByBean(refType reflect.Type, alias string) (columnSql string) {
	columnSql = ""
	fieldCount := refType.NumField() // field count
	for i := 0; i < fieldCount; i++ {
		fieldType := refType.Field(i)         // field type
		column := fieldType.Tag.Get("column") // field tag
		if column != "" {
			if alias != "" {
				columnSql += alias + "."
			}
			columnSql += column + ","
		}
	}
	columnSql = columnSql[0 : len(columnSql)-1]
	return
}

func ResultToBean(rows *sql.Rows, refType reflect.Type) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	fieldCount := refType.NumField() // field count
	jsonKeys := []string{}
	values := []interface{}{}
	valueTypes := []string{}
	for _, column := range columns {
		for i := 0; i < fieldCount; i++ {
			fieldType := refType.Field(i)    // field type
			c := fieldType.Tag.Get("column") // field tag
			if c != column {
				continue
			}
			json := fieldType.Tag.Get("json") // field tag
			jsonKey := json
			if strings.Index(jsonKey, ",") > 0 {
				jsonKey = jsonKey[0:strings.Index(jsonKey, ",")]
			}
			jsonKeys = append(jsonKeys, jsonKey)
			switch fieldType.Type.Name() {
			case "string":
				valueTypes = append(valueTypes, "string")
				var value sql.NullString
				values = append(values, &value)
			case "bool":
				valueTypes = append(valueTypes, "bool")
				var value sql.NullBool
				values = append(values, &value)
			default:
				valueTypes = append(valueTypes, "int")
				var value sql.NullInt64
				values = append(values, &value)
			}
		}
	}
	list := []map[string]interface{}{}
	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		one := make(map[string]interface{})
		for index, jsonKey := range jsonKeys {
			v := values[index]
			switch valueTypes[index] {
			case "string":
				value := v.(*sql.NullString)
				if value.Valid {
					one[jsonKey] = value.String
				} else {
					one[jsonKey] = ""
				}
			case "bool":
				value := v.(*sql.NullBool)
				if value.Valid {
					one[jsonKey] = value.Bool
				} else {
					one[jsonKey] = ""
				}
			default:
				value := v.(*sql.NullInt64)
				if value.Valid {
					one[jsonKey] = value.Int64
				} else {
					one[jsonKey] = ""
				}
			}
		}
		list = append(list, one)
	}
	return list, err
}

func ColumnIsString(columnType *sql.ColumnType) bool {

	scanType := columnType.ScanType()

	fmt.Println(columnType)
	fmt.Println(columnType.Name())

	fmt.Println(scanType)
	fmt.Println(scanType.Name())

	return scanType.Name() == "string"
}

func ColumnIsBool(columnType *sql.ColumnType) bool {

	scanType := columnType.ScanType()

	return scanType.Name() == "bool"
}
func ResultToMap(rows *sql.Rows) ([]map[string][]byte, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	values := []interface{}{}

	for _, _ = range columnTypes {
		var value sql.RawBytes
		values = append(values, &value)
	}
	list := []map[string][]byte{}
	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		one := make(map[string][]byte)

		for index, column := range columns {
			v := values[index]
			value := v.(*sql.RawBytes)
			if value != nil {
				one[column] = (*value)
			} else {
				one[column] = []byte{}
			}
		}

		list = append(list, one)
	}
	return list, err
}
