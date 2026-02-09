package im_bot_test

import (
	"chatservice/pkg/utils"
	"fmt"
	"log"
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
Done! Congratulations on your new bot. You will find it at t.me/ddyy86_bot.
You can now add a description, about section and profile picture for your bot,
see /help for a list of commands. By the way, when you've finished creating your cool bot,
ping our Bot Support if you want a better username for it. Just make sure the bot is fully operational before you do this.

Use this token to access the HTTP API:
5990213484:AAFgyypd0GKNsig465LMJM_-rGj_Itapquw
Keep your token secure and store it safely, it can be used by anyone to control your bot.

For a description of the Bot API, see this page: https://core.telegram.org/bots/api
*/

const (
	testToken       = "5990213484:AAFgyypd0GKNsig465LMJM_-rGj_Itapquw"
	chatID    int64 = -753735474
)

func CreateBot(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	return bot
}

func TestDirectMessage(t *testing.T) {
	bot := CreateBot(testToken)

	sendMessage := "99999999988"
	if bot != nil {
		msg := tgbotapi.NewMessage(chatID, sendMessage)
		bot.Send(msg)
	}
}

func TestCreateIMParam(t *testing.T) {

	type tg struct {
		Token   string `json:"token"`
		ChatId  int64  `json:"chat_id"`
		MsgData string `json:"msg_data"`
	}

	var tgData tg

	tgData.Token = testToken
	tgData.ChatId = chatID
	tgData.MsgData = fmt.Sprintf("現在時間為 %v", time.Now())

	sendJson := utils.ToJSON(tgData)

	sendEncodeStr := utils.EncodeBase64([]byte(sendJson))

	t.Log(sendEncodeStr)

}
