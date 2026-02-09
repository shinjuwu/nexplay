package api

import (
	"chatservice/pkg/utils"
	"database/sql"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func createTelegramBot(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	return bot
}

func TelegramSend(db *sql.DB, payload string) error {

	jsonBytes, err := utils.DecodeBase64(payload)
	if err != nil {
		return err
	}

	paramMap := utils.ToMap(jsonBytes)
	token := utils.ToString(paramMap["token"], "")
	chatID, _ := paramMap["chat_id"].(float64)
	msgData := utils.ToString(paramMap["msg_data"], "")
	if token == "" || chatID == 0 || msgData == "" {
		return fmt.Errorf("parse param has error")
	}

	bot := createTelegramBot(token)
	if bot != nil {
		msg := tgbotapi.NewMessage(int64(chatID), msgData)
		if _, err := bot.Send(msg); err != nil {
			return err
		}

	}

	return nil
}
