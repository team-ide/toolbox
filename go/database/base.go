package database_

import (
	"database/sql"
)

type SqlParam struct {
	Sql    string        `json:"sql,omitempty"`
	Params []interface{} `json:"params,omitempty"`
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
