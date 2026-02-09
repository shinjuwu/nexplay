package api_module_test

import (
	"monitorservice/pkg/utils"
	"testing"
)

func TestXxx(t *testing.T) {

	res, err := utils.PostAPI("http://172.30.0.150:9986/api/v1/system/getserversetting", "application/json", "", "")
	if err != nil {
		t.Logf("error is %v", err)
	} else {
		t.Log(res)
	}

}
