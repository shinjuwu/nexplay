package system

func createBackendActionLogWithTitle(title interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["title"] = title
	return m
}

func createBancendActionLogDetail(before interface{}, after interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	m["before"] = before
	m["after"] = after
	return m
}
