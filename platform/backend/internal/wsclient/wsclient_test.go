package wsclient_test

import (
	"backend/internal/wsclient"
	"testing"
)

func TestGetLoginToken(t *testing.T) {

	mm := wsclient.NewChatServiceManager()

	connInfoStr := `{"scheme":"http","host":"127.0.0.1","port":"8986","domain":"127.0.0.1:8986","api_key":"defaultkey","api_Login_path":"/chatservice.api/v1/login","ws_conn_path":"/chatservice.ws"}`

	mm.ParseConnInfo(connInfoStr)

	token := mm.Login("gg1234", 99)
	t.Logf("token: %s", token)
}

func Benchmark_Division(b *testing.B) {
	mm := wsclient.NewChatServiceManager()

	connInfoStr := `{"scheme":"http","host":"127.0.0.1","port":"8986","domain":"127.0.0.1:8986","api_key":"defaultkey","api_Login_path":"/chatservice.api/v1/login","ws_conn_path":"/chatservice.ws"}`

	mm.ParseConnInfo(connInfoStr)

	for i := 0; i < b.N; i++ {
		token := mm.Login("gg1234", 99)
		b.Logf("token: %s", token)
	}
}
