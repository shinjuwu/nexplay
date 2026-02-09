package utils_test

import (
	"backend/pkg/utils"
	"log"
	"net/url"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	addr := "http://172.30.0.154:9642/deposit"
	// type Deposit struct {
	// 	Method string  `json:"method"`
	// 	Gold   float64 `json:"gold"`
	// 	UserId int64   `json:"userid"`
	// }

	// temp := &Deposit{
	// 	Method: "up",
	// 	Gold:   100.0,
	// 	UserId: 80012,
	// }
	form := url.Values{}

	var gold int64 = 100
	userId := 49
	form.Add("method", "up")
	form.Add("gold", strconv.FormatInt(gold, 10))
	form.Add("userid", strconv.Itoa(userId))

	log.Printf("form value: %v", form)

	resultStr, err := utils.PostAPI(addr, "application/x-www-form-urlencoded", "", form.Encode())
	if err != nil {
		log.Println(err)
	}
	log.Println(resultStr)
}

func Test123(t *testing.T) {
	// addr := "https://egp.qa168.top/86Game/DoAction"

	/*
			  "path": "/86Game/DoAction",
		  "domain": "egp.qa168.top",
		  "scheme": "https",
		  "api_key": ""
	*/
	u := url.URL{
		Scheme: "https",
		Host:   "egp.qa168.top",
		Path:   "/86Game/DoAction"}
	// RawQuery: rawApiQuery.Encode()}

	// log.Printf("form value: %v", form)

	resultStr, httpCode, err := utils.GetBasicAuthAPIWithHttpCode(u.String(), "", "")
	if err != nil {
		log.Println(err)
	}
	log.Println(resultStr)
	log.Println(httpCode)
}
