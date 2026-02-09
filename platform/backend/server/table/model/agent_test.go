package model_test

import (
	"backend/pkg/utils"
	"backend/server/table/model"
	"testing"
)

func TestWalletConnInfo(t *testing.T) {

	ww := new(model.WalletConnInfo)

	urlString := "https://www.youtube.com/watch?v=GHXr4bBxHCo&ab_channel=gemVEVO"

	if success, _ := ww.ParseUrl(urlString); success {
		t.Logf("%v", utils.ToJSON(ww))
	} else {
		t.Logf("%v", ww)
	}

}
