package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/goinggo/mapstructure"
)

// Converter utilites

// ToJSON convert interface{} -> json
func ToJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(js)
}

// ToMap convert json -> map[string]interface{}
func ToMap(b []byte) map[string]interface{} {
	var err error
	var jsonMap = make(map[string]interface{})
	err = json.Unmarshal(b, &jsonMap)
	if err != nil {
		return nil
	}

	return jsonMap
}

func ToArrayMap(b []byte) []map[string]interface{} {
	var err error
	var jsonMap []map[string]interface{}
	err = json.Unmarshal(b, &jsonMap)
	if err != nil {
		return nil
	}

	return jsonMap
}

func ToStruct(b []byte, result interface{}) error {
	jsonMap := ToMap(b)
	return mapstructure.Decode(jsonMap, result)
}

func StructToMap(in interface{}) (map[string]interface{}, error) {

	var outPutMap map[string]interface{}
	inrec, err := json.Marshal(in)
	if err != nil {
		return outPutMap, err
	}
	err = json.Unmarshal(inrec, &outPutMap)
	return outPutMap, err
}

// ToMapMap convert json -> map[string]map[string]interface{}
func ToMapMap(b []byte) map[string]map[string]interface{} {
	var err error
	var jsonMap = make(map[string]map[string]interface{})
	err = json.Unmarshal(b, &jsonMap)
	if err != nil {
		return nil
	}

	return jsonMap
}

// ToInterface convert json -> interface{}
func ToInterface(b []byte) interface{} {
	var err error
	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil
	}
	return data
}

// ToInt convert interface{} -> int, if failed return default value
func ToInt(m interface{}, defaultValue ...int) int {
	val, err := strconv.ParseInt(fmt.Sprintf("%v", m), 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return int(val)
}

// ToInt32 convert interface{} -> int32, if failed return default value
func ToInt32(m interface{}, defaultValue ...int32) int32 {
	val, err := strconv.ParseInt(fmt.Sprintf("%v", m), 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return int32(val)
}

// ToInt64 convert interface{} -> int64, if failed return default value
func ToInt64(m interface{}, defaultValue ...int64) int64 {
	val, err := strconv.ParseInt(fmt.Sprintf("%v", m), 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	return val
}

// ToString convert interface{} -> string, if failed return default value
func ToString(m interface{}, defaultValue ...string) string {
	if val, ok := m.(string); ok {
		return val
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// ToString convert interface{} -> string, if failed return default value
func ToBool(m interface{}, defaultValue ...bool) bool {
	if val, ok := m.(bool); ok {
		return val
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

func ToTimeUnixUTC(timestamp string, defaultValue ...int64) time.Time {
	if len(timestamp) != 10 {
		if len(defaultValue) > 0 {
			return time.Unix(0, 0).UTC()
		}

		return time.Now().UTC()
	}

	i := ToInt64(timestamp, 0)

	return time.Unix(i, 0).UTC()
}

func ToTimeUnixMilliUTC(timestamp string, defaultValue ...int64) time.Time {
	if len(timestamp) != 13 {
		if len(defaultValue) > 0 {
			return time.UnixMilli(defaultValue[0]).UTC()
		}

		return time.Now().UTC()
	}

	i := ToInt64(timestamp, 0)

	return time.UnixMilli(i).UTC()
}

func ToTimeUnixMicroUTC(timestamp string, defaultValue ...int64) time.Time {
	if len(timestamp) != 16 {
		if len(defaultValue) > 0 {
			return time.UnixMicro(defaultValue[0]).UTC()
		}

		return time.Now().UTC()
	}

	i := ToInt64(timestamp, 0)

	return time.UnixMicro(i).UTC()
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

// hex
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt(s string, defaultValue ...int) int {
	if i, err := strconv.Atoi(s); err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
	} else {
		return i
	}

	return 0
}

// hex
func StringToInt64(s string, defaultValue ...int64) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
	} else {
		return i
	}

	return 0
}
