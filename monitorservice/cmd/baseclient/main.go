package main

import (
	"monitorservice/cmd/baseclient/client"
	"monitorservice/pkg/utils"

	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// var addr = flag.String("addr", "127.0.0.1:8986", "local http service address")

// var addr = flag.String("addr", "103.122.226.37:7350", "uat http service address")
// var addr = flag.String("addr", "172.26.11.164:7350", "dev http service address")

var (
	host         = "127.0.0.1"
	port         = 8986
	domain       = fmt.Sprintf("%s:%d", host, port)
	apiLoginAddr = "/monitorservice.api/v1/login"
	wsAddr       = "/monitorservice.ws"
)

type Message struct {
	Id      string `json:"id"`      // method id (rpc id)
	Subject string `json:"subject"` // sub method id
	Payload string `json:"payload"` // send data
}

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// get login token
	rawApiQuery := make(url.Values)
	rawApiQuery.Add("username", "gg1234")
	rawApiQuery.Add("agent_id", "99")

	u := url.URL{Scheme: "http", Host: domain, Path: apiLoginAddr, RawQuery: rawApiQuery.Encode()}

	body, err := utils.GetAPI(u.String(), "defaultkey", "")
	if err != nil {
		log.Println(err)
	}

	bodyMap := make(map[string]interface{}, 0)
	_ = json.Unmarshal([]byte(body), &bodyMap)

	// log.Println(bodyMap["data"])
	data := bodyMap["data"].(map[string]interface{})
	if utils.ToInt(bodyMap["code"], -1) != 0 {
		fmt.Println("error api reponse")
		fmt.Println(bodyMap["code"])
		return
	}
	token := utils.ToString(data["token"], "")
	if token == "" {
		fmt.Println("token is empty")
		return
	}

	// use token to login ws
	rawQuery := make(url.Values)
	rawQuery.Add("token", token)

	client, err := client.NewWebSocketClient(domain, wsAddr, rawQuery.Encode())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting")

	go func() {
		// write down data every 100 ms
		ticker := time.NewTicker(time.Millisecond * 1500)
		i := 0
		for range ticker.C {
			var m Message
			m.Id = "test"
			m.Subject = "message"
			m.Payload = `{"gameId":1,"roomType":1}`
			err = client.Write(m)
			if err != nil {
				fmt.Printf("error: %v, writing error\n", err)
			}
			i++
		}
	}()

	// Close connection correctly on exit
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// The program will wait here until it gets the
	<-sigs
	client.Stop()
	fmt.Println("Goodbye")

}
