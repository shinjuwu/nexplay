package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

var EmptyBytes = []byte("")

//PostAPI 呼叫Post API
func PostAPI(addr string, contentType, authorization, data string) (result string, err error) {

	var req *http.Request

	if contentType == "application/json" {
		jsonObject := []byte(data)
		req, err = http.NewRequest("POST", addr, bytes.NewBuffer(jsonObject))
	} else if contentType == "application/x-www-form-urlencoded" {

		req, err = http.NewRequest("POST", addr, strings.NewReader(data))
	}

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("authorization", authorization)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	result = string(body)
	return
}

//GetAPI 呼叫Get API
func GetAPI(urlStr, basicAuthUsername, basicAuthPassword string) ([]byte, error) {
	req, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		return EmptyBytes, err
	}

	if basicAuthUsername != "" || basicAuthPassword != "" {
		req.SetBasicAuth(basicAuthUsername, basicAuthPassword)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return EmptyBytes, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
