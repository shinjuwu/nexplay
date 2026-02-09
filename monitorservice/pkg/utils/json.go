package utils

import "encoding/json"

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
