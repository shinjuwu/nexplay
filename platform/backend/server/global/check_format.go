package global

import "strings"

/*
check order id format
格式:代理編號+yyyyMMddHHmmssSSS+account
*/
func CheckOrderIdFormat(orderId, agentId, username string) (errMsg string, isSuccess bool) {
	isSuccess = false
	if len(orderId) > 100 {
		errMsg = "order id length is illegal"
		return
	}

	agentIdLength := len(agentId)

	if strings.Compare(agentId, orderId[0:agentIdLength]) != 0 {
		errMsg = "order id format1 is illegal"
		return
	}

	if strings.Compare(username, orderId[agentIdLength+17:]) != 0 {
		errMsg = "order id format2 is illegal"
		return
	}

	isSuccess = true

	return
}
