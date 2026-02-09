package utils

import (
	"reflect"
	"strings"
)

/*
取得指定結構 tag 名稱的列表
*/
func GetFieldsName(tag string, s interface{}) (fieldnames []string) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get(tag), ",")[0] // use split to ignore tag "options" like omitempty, etc.
		fieldnames = append(fieldnames, v)
	}
	return
}
