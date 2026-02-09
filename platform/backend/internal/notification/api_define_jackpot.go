package notification

import (
	"backend/pkg/utils"
	"fmt"
)

const (
	ApiNotification_jppool = "jppool" // jackpot獎池資料
	ApiNotification_crank  = "crank"  // jackpot玩家貢獻度
	ApiNotification_jpon   = "jpon"   // jackpot總開關開啟
	ApiNotification_jpoff  = "jpoff"  // jackpot總開關關閉

)

func GetJpPool() (float64, float64, float64, float64, float64, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_jppool

	resultStr, err := utils.GetAPI(addr, "", "")
	if err != nil {
		fmt.Printf("GetJpPool response: %s, err: %v", resultStr, err)
		return .0, .0, .0, .0, .0, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))
	// {"code":0,"message":"success","data":"{\"JPShowPool\":10000,\"JPinjectWaterRate\":0.001,\"RealPool\":0.005,\"ReservePool\":0,\"ShowPool\":10000}"}

	code, ok := result["code"].(float64)

	if !ok {
		return .0, .0, .0, .0, .0, API_OTHER_ERROR, fmt.Errorf("GetJpPool has error, response body is: %v", resultStr)
	}

	if code != 0 {
		return .0, .0, .0, .0, .0, utils.ToInt(code), fmt.Errorf("GetJpPool has error, code is: %v", code)
	}

	dataStr, ok := result["data"].(string)
	if !ok {
		return .0, .0, .0, .0, .0, utils.ToInt(code), fmt.Errorf("GetJpPool has error, response body is: %v", resultStr)
	}
	data := utils.ToMap([]byte(dataStr))

	realPool := float64(0)
	showPool := float64(0)
	reservePool := float64(0)

	jpInjectWaterRate := float64(0)
	jpShowPool := float64(0)

	if pool, ok := data["RealPool"].(float64); ok {
		realPool = pool
	}
	if pool, ok := data["ShowPool"].(float64); ok {
		showPool = pool
	}
	if pool, ok := data["ReservePool"].(float64); ok {
		reservePool = pool
	}
	if rate, ok := data["JPinjectWaterRate"].(float64); ok {
		jpInjectWaterRate = rate
	}
	if pool, ok := data["JPShowPool"].(float64); ok {
		jpShowPool = pool
	}

	return realPool, showPool, reservePool, jpShowPool, jpInjectWaterRate, utils.ToInt(code), nil
}

func GetCRank() ([]map[string]interface{}, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_crank

	resultStr, err := utils.GetAPI(addr, "", "")
	if err != nil {
		fmt.Printf("GetCRank response: %s, err: %v", resultStr, err)
		return nil, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))
	// {"code":0,"message":"success","data":"[{\"UserID\":\"1083\",\"PlayNum\":6,\"TotalBet\":51,\"Win\":6,\"Lose\":0,\"Score\":0.11764705882352941},{\"UserID\":\"81496712\",\"PlayNum\":1,\"TotalBet\":3,\"Win\":0,\"Lose\":1,\"Score\":0},{\"UserID\":\"2693\",\"PlayNum\":6,\"TotalBet\":0,\"Win\":0,\"Lose\":6,\"Score\":0},{\"UserID\":\"1001\",\"PlayNum\":1,\"TotalBet\":1,\"Win\":0,\"Lose\":1,\"Score\":0}]"}

	code, ok := result["code"].(float64)

	if !ok {
		return nil, API_OTHER_ERROR, fmt.Errorf("GetCRank has error, response body is: %v", resultStr)
	}

	if code != 0 {
		return nil, utils.ToInt(code), fmt.Errorf("GetCRank has error, code is: %v", code)
	}

	dataStr, ok := result["data"].(string)
	if !ok {
		return nil, utils.ToInt(code), fmt.Errorf("GetCRank has error, response body is: %v", resultStr)
	}
	data := utils.ToArrayMap([]byte(dataStr))

	return data, utils.ToInt(code), nil
}

func GetJpOn() (int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_jpon

	resultStr, err := utils.GetAPI(addr, "", "")
	if err != nil {
		fmt.Printf("GetJpOn response: %s, err: %v", resultStr, err)
		return API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))
	// {"code":0,"message":"success"}

	code, ok := result["code"].(float64)

	if !ok {
		return API_OTHER_ERROR, fmt.Errorf("GetJpOn has error, response body is: %v", resultStr)
	}

	if code != 0 {
		return utils.ToInt(code), fmt.Errorf("GetCRank has error, code is: %v", code)
	}

	return utils.ToInt(code), nil
}

func GetJpOff() (int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_jpoff

	resultStr, err := utils.GetAPI(addr, "", "")
	if err != nil {
		fmt.Printf("GetJpOff response: %s, err: %v", resultStr, err)
		return API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))
	// {"code":0,"message":"success"}

	code, ok := result["code"].(float64)

	if !ok {
		return API_OTHER_ERROR, fmt.Errorf("GetJpOff has error, response body is: %v", resultStr)
	}

	if code != 0 {
		return utils.ToInt(code), fmt.Errorf("GetJpOff has error, code is: %v", code)
	}

	return utils.ToInt(code), nil
}
