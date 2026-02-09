package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
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
func GetAPI(url string, parameter string, authorization string) (result string, err error) {
	urlStr := fmt.Sprintf(`%s?%s`, url, parameter)

	req, err := http.NewRequest("GET", urlStr, nil)
	req.Header.Set("authorization", authorization)
	if err != nil {
		return
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	result = string(body)
	return
}

func GetBasicAuthAPIWithHttpCode(urlStr, basicAuthUsername, basicAuthPassword string) ([]byte, int, error) {
	req, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		return EmptyBytes, http.StatusBadRequest, err
	}

	if basicAuthUsername != "" || basicAuthPassword != "" {
		req.SetBasicAuth(basicAuthUsername, basicAuthPassword)
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		return EmptyBytes, http.StatusBadRequest, err
	}
	defer response.Body.Close()

	bs, err := ioutil.ReadAll(response.Body)

	return bs, response.StatusCode, err
}

//GetAPI 呼叫Get API
func GetBasicAuthAPI(urlStr, basicAuthUsername, basicAuthPassword string) ([]byte, error) {
	req, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		return EmptyBytes, err
	}

	if basicAuthUsername != "" || basicAuthPassword != "" {
		req.SetBasicAuth(basicAuthUsername, basicAuthPassword)
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		return EmptyBytes, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

//GetAPI 呼叫Get API
func PostBasicAuthAPI(urlStr, basicAuthUsername, basicAuthPassword, data string) ([]byte, error) {

	var req *http.Request

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return EmptyBytes, err
	}

	if basicAuthUsername != "" {
		req.SetBasicAuth(basicAuthUsername, basicAuthPassword)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Do(req)
	if err != nil {
		return EmptyBytes, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
