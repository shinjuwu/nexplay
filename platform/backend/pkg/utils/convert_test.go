package utils_test

import (
	"backend/api/intercom/model"
	"backend/pkg/utils"
	"log"
	"testing"
)

func TestConvertTime(t *testing.T) {
	// 2017-03-06 14:30:36 +0800 CST
	timeString := "1488781836"
	temp := utils.ToTimeUnixUTC(timeString, 0)
	log.Println(temp)

	timeMicroString := "1488781836949856"
	temp = utils.ToTimeUnixMicroUTC(timeMicroString, 0)
	log.Println(temp)

	// 2017-03-06 14:30:36.949 +0800 CST
	timeMilliString := "1488781836949"
	temp = utils.ToTimeUnixMilliUTC(timeMilliString, 0)
	log.Println(temp)

	// 1970-01-01 00:00:00 +0000 UTC
	timeErrorString := "123"
	temp = utils.ToTimeUnixUTC(timeErrorString, 0)
	log.Println(temp)

}

func TestToStruct(t *testing.T) {
	resp := model.GetGameServerStateResponse{}

	jsonStr := `{"state": 1}`
	err := utils.ToStruct([]byte(jsonStr), &resp)
	if err != nil {
		return
	}

	log.Println(resp)
}
