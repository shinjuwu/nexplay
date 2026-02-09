package global

import (
	"backend/internal/notification"
	"backend/pkg/utils"
	"backend/server/table/model"
	"definition"
)

func IMAlertTelegram(imNotifys []string) {

	conninfo := make(map[string]interface{}, 0)
	addr, ok := ServerInfoCache.Load("chat")
	if ok {
		addres, ok := addr.(model.ServerInfo)
		if ok {
			conninfo = utils.ToMap(addres.AddressesBytes)
		}
	}

	type tg struct {
		Token   string `json:"token"`
		ChatId  int64  `json:"chat_id" mapstructure:"chat_id"`
		MsgData string `json:"msg_data"`
	}

	storage, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_IM_ALERT_TG)
	if !ok {
		return
	}

	var tgData tg

	err := utils.ToStruct([]byte(storage.Value), &tgData)
	if err != nil {
		return
	}

	tgData.MsgData = "本次被自動風控的用戶資訊如下：" + utils.ToJSON(imNotifys)

	notification.SendNotifyToIM(conninfo, definition.IM_TYPE_TELEGRAM, utils.ToJSON(tgData))
	// add im notify end.
}
