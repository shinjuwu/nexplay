package database

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"time"
)

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func QueryAndAutoParse(db *sql.DB, query string, args ...interface{}) (response []map[string]interface{}, err error) {
	if rows, err := db.Query(query, args...); err == nil {
		defer rows.Close()
		if columns, err := rows.Columns(); err == nil {
			for rows.Next() {
				values := make([]interface{}, len(columns))
				pointers := make([]interface{}, len(columns))
				for i := range values {
					pointers[i] = &values[i]
				}
				_ = rows.Scan(pointers...)
				resultMap := make(map[string]interface{})
				for i, val := range values {
					var covert interface{}
					switch x := val.(type) {
					case int:
						covert = val.(int)
					case float64:
						covert = val.(float64)
					case []uint8:
						covert = B2S(val.([]uint8))
					case bool:
						covert = val.(bool)
					case string:
						covert = val
					case int64:
						covert = val.(int64)
					case time.Time:
						covert = val
					case nil:
						covert = val
					default:
						log.Printf("unknow type: %v, value: %v", reflect.TypeOf(x), x)
						covert = val
					}
					resultMap[columns[i]] = covert
				}
				response = append(response, resultMap)
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
	return
}

func RowsTransformStruct(rows *sql.Rows, dest interface{}, structTag string) error {
	col_names, err := rows.Columns()
	if err != nil {
		return err
	}

	v := reflect.ValueOf(dest)
	if v.Elem().Type().Kind() != reflect.Struct {
		return errors.New("target not a struct")
	}

	scan_dest := []interface{}{}

	addr_by_col_name := map[string]interface{}{}

	for i := 0; i < v.Elem().NumField(); i++ {
		propertyName := v.Elem().Field(i)
		col_name := v.Elem().Type().Field(i).Tag.Get(structTag)
		if col_name == "" {
			if !v.Elem().Field(i).CanInterface() {
				continue
			}
			col_name = propertyName.Type().Name()
		}

		addr_by_col_name[col_name] = propertyName.Addr().Interface()
	}

	for _, col_name := range col_names {
		scan_dest = append(scan_dest, addr_by_col_name[col_name])
	}

	return rows.Scan(scan_dest...)
}

func DoQuery(db *sql.DB, sqlInfo string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(sqlInfo, args...)
	if err != nil {
		return nil, err
	}
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	for index := range cache {                 //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}
	_ = rows.Close()
	return list, nil
}
