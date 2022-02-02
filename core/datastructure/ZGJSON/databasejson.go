package ZGJSON

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

func JsonFromDbRows(rows *sql.Rows) []string {
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	c := 0
	results := make(map[string]interface{})
	data := []string{}

	for rows.Next() {
		if c > 0 {
			data = append(data, ",")
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for i, value := range values {
			switch value.(type) {
			case nil:
				results[columns[i]] = nil

			case []byte:
				s := string(value.([]byte))
				x, err := strconv.Atoi(s)

				if err != nil {
					results[columns[i]] = s
				} else {
					results[columns[i]] = x
				}

			default:
				results[columns[i]] = value
			}
		}

		b, _ := json.Marshal(results)
		data = append(data, strings.TrimSpace(string(b)))
		c++
	}

	return data
}
func JsonFromDriverRows(rowsc driver.Rows) []string {
	var columns = rowsc.(driver.RowsColumnTypeScanType).Columns()
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	c := 0
	results := make(map[string]interface{})
	nowrecorddata := make([]driver.Value, len(columns))
	data := []string{}
	for {
		if err := rowsc.Next(nowrecorddata); err != nil {
			if err == io.EOF {
				break
			}
			rowsc.Close()
		} else { // OK data

			for i, value := range nowrecorddata {
				switch value.(type) {
				case nil:
					results[columns[i]] = nil

				case []byte:
					s := string(value.([]byte))
					x, err := strconv.Atoi(s)

					if err != nil {
						results[columns[i]] = s
					} else {
						results[columns[i]] = x
					}

				default:
					results[columns[i]] = value
				}
			}
			b, _ := json.Marshal(results)
			data = append(data, strings.TrimSpace(string(b)))
			c++
		}
	}
	return data
}
