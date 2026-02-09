package utils

import (
	"sort"
	"strings"
	"time"
)

func TimeNowUTC() time.Time {
	return time.Now().UTC()
}

/*
get unsigned time format string.

example: 2006-01-02 15:04:05
*/
func GetTimeNowUTC() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

/*
get unsigned time format string.

example: 20060102150405
*/
func GetUnsignedTimeNowUTC() string {
	return time.Now().UTC().Format("200601021504")
}

/*
get unsigned time format string.

example: 20060102150405
*/
func GetUnsignedTimeUTC(t time.Time, format string) string {
	switch format {
	case "15min":
		fallthrough
	case "minute":
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

/*
get unsigned time format string.

example: 20060102150405
*/
func GetUnsignedTimeUTCFromStr(s string, format string) (time.Time, error) {
	switch format {
	case "15min":
		fallthrough
	case "minute":
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

	return time.ParseInLocation(format, s, time.UTC)
}

/*
get unix time.

example: 1640860964
*/
func GetUnixTimeNowUTC() int64 {
	return time.Now().UTC().Unix()
}

/*
get unsigned time format string.

example: 200601021504050000
*/
func Get18UnsignedTimeNowUTC() string {
	t := time.Now().UTC().Format("20060102150405.0000") //20220712053750.6464
	t = strings.Replace(t, ".", "", 1)
	return t
}

func GetTimeNowUTCTodayTime() time.Time {
	t := time.Now().UTC()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.UTC().Location())
}

func TruncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.UTC().Location())
}

func TruncateToHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.UTC().Location())
}

func TruncateToMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.UTC().Location())
}

// monday is first day
func TruncateToWeek(t time.Time) time.Time {
	offset := (int(time.Monday) - int(t.Weekday()) - 7) % 7
	result := t.Add(time.Duration(offset*24) * time.Hour)
	return time.Date(result.Year(), result.Month(), result.Day(), 0, 0, 0, 0, t.UTC().Location())
}

// check date is monday
func CheckIsMonday(t time.Time) bool {
	return t.Weekday() == time.Monday
}

/* 依照輸入參數類型，回傳時間區段內的指定時間類型的第一個資料列表
* hour format YYYYMMDDhh

* day format YYYYMMDD

* week format YYYYMMDD : return every Monday

* month format YYYYMM
 */
func GetTimeIntervalList(dateType string, start, end time.Time) []string {

	args := make([]string, 0)

	inerval := int64(0)
	format := ""

	switch dateType {
	case "minute":
		inerval = 86400 / 24 / 60
		format = "200601021504"
	case "15min":
		inerval = 86400 / 24 / 4
		format = "200601021504"
	case "hour":
		inerval = 86400 / 24
		format = "2006010215"
	case "day":
		inerval = 86400
		format = "20060102"
	case "week":
		start = TruncateToWeek(start)
		end = TruncateToWeek(end)
		inerval = (86400 * 7) + 1
		format = "20060102"
	case "month":
		inerval = 86400 * 31
		format = "200601"
	default:
		return args
	}

	sInt := start.Unix()
	eInt := end.Unix()

	args = calAllIntervalList(format, inerval, sInt, eInt)

	return args
}

// 取得指定時間區段內所有的日期列表
func GetAllIntervalList(dateType string, start time.Time) (args []string) {
	inerval := int64(0)
	format := ""

	var sInt, eInt int64

	switch dateType {
	// case "hour":
	// 	inerval = 86400 / 24
	// 	format = "2006010215"
	// 	start = TruncateToHour(start)
	// 	end := start.AddDate(0, 0, 1)
	// 	sInt = start.Unix()
	// 	eInt = end.Unix() - 1
	case "day":
		inerval = 86400
		format = "20060102"
		end := start.AddDate(0, 0, 1)
		sInt = start.Unix()
		eInt = end.Unix() - 1
	case "week":
		start = TruncateToWeek(start)
		inerval = 86400
		format = "20060102"
		sInt = start.Unix()
		eInt = sInt + (inerval * 6)
	case "month":
		start = TruncateToMonth(start)
		end := start.AddDate(0, 1, 0)
		inerval = 86400
		format = "20060102"
		sInt = start.Unix()
		eInt = end.Unix() - 1
	default:
		return args
	}

	return calAllIntervalList(format, inerval, sInt, eInt)
}

func calAllIntervalList(format string, inerval, sInt, eInt int64) []string {

	args := make([]string, 0)
	tmps := make(map[string]string, 0)

	if sInt < eInt {
		pst := time.Unix(sInt, 0).UTC().Format(format)
		pet := time.Unix(eInt, 0).UTC().Format(format)
		tmps[pst] = pst
		tmps[pet] = pet
		for {
			st := time.Unix(sInt, 0).UTC().Format(format)
			tmps[st] = st
			sInt += inerval
			if sInt > eInt {
				// log.Printf("length: %v", len(args))
				for timeStr := range tmps {
					args = append(args, timeStr)
				}
				sort.Strings(args)
				return args
			} else {
				tmps[st] = st
			}
		}
	}

	return args
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

func MicroSecondsToUtcTime(usec int64) time.Time {
	return time.UnixMicro(usec).UTC()
}

func GetTimeUTCToday(offset int) (time.Time, time.Time) {

	transHour := offset / 60
	transMin := offset % 60

	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month(), now.Day(), transHour, transMin, 0, 0, time.UTC)
	end := time.Date(now.Year(), now.Month(), now.Day()+1, transHour, transMin, 0, 0, time.UTC)

	// 輸出結果
	// fmt.Printf("UTC %+d 的開始時間為 %s，結束時間為 %s\n", offset, start.String(), end.String())

	return start, end
}

func GetTimeUTCThisMonth(offset int) (time.Time, time.Time) {

	transHour := offset / 60
	transMin := offset % 60

	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month(), 1, transHour, transMin, 0, 0, time.UTC)
	end := time.Date(now.Year(), now.Month()+1, 1, transHour, transMin, 0, 0, time.UTC)

	return start, end
}

func GetTimeUTCThisWeek(offset int) (time.Time, time.Time) {

	transHour := offset / 60
	transMin := offset % 60

	now := TruncateToWeek(time.Now().UTC())
	start := time.Date(now.Year(), now.Month(), now.Day(), transHour, transMin, 0, 0, time.UTC)
	end := time.Date(now.Year(), now.Month(), now.Day()+7, transHour, transMin, 0, 0, time.UTC)

	return start, end
}

func Get15MinFormatString(t time.Time) string {
	/*
		>= 2023-01-01T12:00:00Z && < 2023-01-01T12:15:00Z -> 202301011200
		>= 2023-01-01T12:15:00Z && < 2023-01-01T12:30:00Z -> 202301011215
	*/
	roundedMinutes := (t.Minute() / 15) * 15
	if roundedMinutes == 60 {
		roundedMinutes = 0
		t = t.Add(time.Hour) // 增加一小時
	}
	result := time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), roundedMinutes, 0, 0, t.Location(),
	)

	return TransUnsignedTimeUTCFormat("15min", result)
}
