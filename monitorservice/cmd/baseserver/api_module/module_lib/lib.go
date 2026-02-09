package module_lib

import (
	"fmt"
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/internal/api"
	"monitorservice/pkg/jwt"
	"monitorservice/pkg/utils"
	"time"
)

// 指定某個用戶斷線
// api token 失效
func LibBlockUsers(jwt *jwt.JwtManager, username string) {

	jwt.AddBlackTokenByUsername(username)
}

// 呼叫 API 確認服務目前狀態
func CallAPICheckServiceStatus(d *model.ServiceInfo) int {
	if d.APIURLs.Method == "POST" {
		_, err := utils.PostAPI(d.APIURLs.URL, "application/json", "", "")
		if err != nil {
			return api.ERROR_CODE_FAIL
		}

	} else if d.APIURLs.Method == "GET" {
		resultStr, err := utils.GetAPI(d.APIURLs.URL, d.APIURLs.AuthKey, "")
		if err != nil {
			return api.ERROR_CODE_FAIL
		}

		result := &model.Response{}
		utils.ToStruct([]byte(resultStr), result)

		if result.Code != api.ERROR_CODE_SUCCESS {
			return api.ERROR_CODE_FAIL
		}
	} else {
		return api.ERROR_CODE_FAIL
	}

	return api.ERROR_CODE_SUCCESS
}

// 檢查平台碼是否有效
func CheckPermissionWithClaim(permissions []string, filter string) bool {
	permissionMap := make(map[string]bool)
	for i := 0; i < len(permissions); i++ {
		permissionMap[permissions[i]] = true
	}

	return permissionMap[filter]
}

/*
依照目前時間產生格式化表名稱(tablename_yyyymmdd)
	產生本月相關的資料表名稱
	以查詢 UTC+8 為例，timZoneMin 輸入 -480
	ex : 2023-12-01T00:00:00 ~ 2023-12-31T23:59:59 (UTC+8)
	實際作用的時間區段是 : 2023-11-30T16:00:00 ~ 2023-12-31T15:59:59 (UTC+0)
	產生的資料表名稱陣列是 [202311 202312]
*/
func GenerateTableNameMonth(timZoneMin int) []string {
	transHour := timZoneMin / 60
	transMin := timZoneMin % 60

	nowTime := time.Now().UTC()

	timeThisTimeStart := time.Date(nowTime.Year(), nowTime.Month(), 1, transHour, transMin, 0, 0, nowTime.UTC().Location())
	timeThisTimeEnd := time.Date(nowTime.Year(), nowTime.Month()+1, 1, transHour, transMin, 0, 0, nowTime.UTC().Location())

	res := []string{
		TransUnsignedTimeUTCFormat("month", timeThisTimeStart),
		TransUnsignedTimeUTCFormat("month", timeThisTimeEnd),
	}

	return res
}

func GenerateMonthPeriodOfTime(outputFormat string, timZoneMin int) []string {
	transHour := timZoneMin / 60
	transMin := timZoneMin % 60

	nowTimeFromTimeZone := time.Now().UTC().Add(time.Duration(-timZoneMin) * time.Minute)
	nowTime := time.Date(nowTimeFromTimeZone.Year(), nowTimeFromTimeZone.Month(), nowTimeFromTimeZone.Day(), 0, 0, 0, 0, nowTimeFromTimeZone.UTC().Location())

	timeThisTimeStart := time.Date(nowTime.Year(), nowTime.Month(), 1, transHour, transMin, 0, 0, nowTime.UTC().Location())
	timeThisTimeEnd := time.Date(nowTime.Year(), nowTime.Month()+1, 1, transHour, transMin, 0, 0, nowTime.UTC().Location())

	res := []string{
		TransUnsignedTimeUTCFormat(outputFormat, timeThisTimeStart),
		TransUnsignedTimeUTCFormat(outputFormat, timeThisTimeEnd),
	}

	return res
}

func GenerateWeekPeriodOfTime(outputFormat string, timZoneMin int) []string {
	transHour := timZoneMin / 60
	transMin := timZoneMin % 60

	nowTimeFromTimeZone := time.Now().UTC().Add(time.Duration(-timZoneMin) * time.Minute)
	nowTimeToDay := time.Date(nowTimeFromTimeZone.Year(), nowTimeFromTimeZone.Month(), nowTimeFromTimeZone.Day(), 0, 0, 0, 0, nowTimeFromTimeZone.UTC().Location())

	nowTime := TruncateToWeek(nowTimeToDay)

	timeThisTimeStart := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), transHour, transMin, 0, 0, nowTime.UTC().Location())
	timeThisTimeEnd := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+7, transHour, transMin, 0, 0, nowTime.UTC().Location())

	res := []string{
		TransUnsignedTimeUTCFormat(outputFormat, timeThisTimeStart),
		TransUnsignedTimeUTCFormat(outputFormat, timeThisTimeEnd),
	}

	return res
}

func GenerateDayPeriodOfTime(outputFormat string, timZoneMin int) []string {
	transHour := timZoneMin / 60
	transMin := timZoneMin % 60

	nowTimeFromTimeZone := time.Now().UTC().Add(time.Duration(-timZoneMin) * time.Minute)
	nowTime := time.Date(nowTimeFromTimeZone.Year(), nowTimeFromTimeZone.Month(), nowTimeFromTimeZone.Day(), 0, 0, 0, 0, nowTimeFromTimeZone.UTC().Location())
	timeThisTimeStart := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), transHour, transMin, 0, 0, nowTime.UTC().Location())
	timeThisTimeEnd := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+1, transHour, transMin, 0, 0, nowTime.UTC().Location())

	res := []string{
		TransUnsignedTimeUTCFormat(outputFormat, timeThisTimeStart),
		TransUnsignedTimeUTCFormat(outputFormat, timeThisTimeEnd),
	}

	return res
}

// 依照目前時間產生格式化字串(yyyymm)
func GenerateTimeStrMonth() string {
	now := time.Now().UTC()
	year := now.Year()
	month := now.Month()
	return fmt.Sprintf("%04d%02d", year, month)
}

// 依照目前時間產生格式化字串(yyyymmdd)
func GenerateTimeStrDay() string {
	now := time.Now().UTC()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	return fmt.Sprintf("%04d%02d%02d", year, month, day)
}

// 依照目前時間產生格式化字串(yyyymmdd)
// The first day is Monday
func GenerateTimeStrWeek() map[string]bool {
	weekStr := make(map[string]bool, 0)
	today := time.Now()
	startOfWeek := today.AddDate(0, 0, -int(today.Weekday())+1)
	for i := 0; i < 7; i++ {
		currentDay := startOfWeek.AddDate(0, 0, i)
		formattedDate := currentDay.Format("20060102")
		weekStr[formattedDate] = true
	}
	return weekStr
}

func TransUnsignedTimeUTCFormat(dateType string, t time.Time) string {
	format := ""
	switch dateType {
	case "minute":
		fallthrough
	case "15min":
		format = "200601021504"
	case "hour":
		format = "2006010215"
	case "day":
		fallthrough
	case "week":
		format = "20060102"
	case "month":
		format = "200601"
	}
	return t.UTC().Format(format)
}

// monday is first day
func TruncateToWeek(t time.Time) time.Time {
	offset := (int(time.Monday) - int(t.Weekday()) - 7) % 7
	result := t.Add(time.Duration(offset*24) * time.Hour)
	return time.Date(result.Year(), result.Month(), result.Day(), 0, 0, 0, 0, t.UTC().Location())
}
