package utils_test

import (
	"backend/pkg/utils"
	"log"
	"strings"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {

	now := time.Now().UTC()

	temp := utils.GetUnsignedTimeNowUTC()
	log.Println(temp)

	temp = utils.GetUnsignedTimeUTC(now, "minute")
	log.Println(temp)
	temp = utils.GetUnsignedTimeUTC(now, "hour")
	log.Println(temp)
	temp = utils.GetUnsignedTimeUTC(now, "day")
	log.Println(temp)
	temp = utils.GetUnsignedTimeUTC(now, "month")
	log.Println(temp)

	temp = time.Now().UTC().Format("20060102150405.0000") //20220712053750.6464
	log.Println(temp)
	temp = strings.Replace(temp, ".", "", 1)
	log.Println(temp)
}

func TestTimeTruncate(t *testing.T) {

	temp := utils.TruncateToHour(time.Now().UTC())
	log.Println(temp)
	temp = utils.TruncateToDay(time.Now().UTC())
	log.Println(temp)
	temp = utils.TruncateToWeek(time.Now().UTC())
	log.Println(temp)
	temp = utils.TruncateToMonth(time.Now().UTC())
	log.Println(temp)

	temp = utils.TruncateToMonth(time.Date(2022, 11, 18, 16, 53, 12, 0, time.Now().UTC().Location()))
	log.Println(temp)

}

func TestGetTimeIntervalList(t *testing.T) {

	temp := utils.GetTimeIntervalList("hour", time.Now().UTC(), time.Now().UTC().Add(25*60*time.Minute))
	log.Printf("hour: %v", temp)

	temp = utils.GetTimeIntervalList("day", time.Now().UTC(), time.Now().UTC().Add(25*60*time.Minute))
	log.Printf("day: %v", temp)

	startTimeTrans := time.Date(2022, 10, 24, 0, 0, 1, 0, time.Now().UTC().Location())
	endTimeTrans := time.Now().UTC()
	temp = utils.GetTimeIntervalList("week", startTimeTrans, endTimeTrans)
	// temp = utils.GetTimeIntervalList("week", time.Now().UTC(), time.Now().UTC().Add(7*24*60*time.Minute))
	log.Printf("week: %v", temp)
	temp = utils.GetTimeIntervalList("month", time.Now().UTC(), time.Now().UTC().Add(10*24*60*time.Minute))
	log.Printf("month: %v", temp)
}

func TestGetWeekIntervalList(t *testing.T) {

	startTimeTrans := time.Date(2022, 10, 24, 0, 0, 1, 0, time.Now().UTC().Location())

	temp := utils.GetAllIntervalList("day", startTimeTrans)
	log.Printf("day: %v", temp)
	// endTimeTrans := time.Now().UTC()
	temp = utils.GetAllIntervalList("week", startTimeTrans)
	log.Printf("week: %v", temp)

	temp = utils.GetAllIntervalList("month", startTimeTrans)
	log.Printf("month: %v", temp)
}

func TestTransUnsignedTimeUTCFormat(t *testing.T) {
	startTime := time.Date(2022, 10, 24, 0, 46, 0, 0, time.Now().UTC().Location())
	// result := startTimeTrans.Truncate(15 * time.Minute)

	rem := startTime.Minute() % 15
	if rem > 0 {
		rem = 1
	} else {
		rem = 0
	}

	roundedMinutes := (startTime.Minute()/15 + rem) * 15
	if roundedMinutes == 60 {
		roundedMinutes = 0
		startTime = startTime.Add(time.Hour) // 增加一小時
	}
	result := time.Date(
		startTime.Year(), startTime.Month(), startTime.Day(),
		startTime.Hour(), roundedMinutes, 0, 0, startTime.Location(),
	)

	temp := utils.TransUnsignedTimeUTCFormat("15min", result)

	log.Printf("15min: %v", temp[0:len("YYYYMM")])
}

func TestGetTimeUTCToday(t *testing.T) {
	// utils.GetTimeUTCToday(-480)

	t.Log(utils.GetTimeUTCToday(-480))
}

func TestGetTimeUTCThisMonth(t *testing.T) {
	// utils.GetTimeUTCToday(-480)
	var timeThisTimeStart, timeThisTimeEnd time.Time
	var timeLastTimeStart, timeLastTimeEnd time.Time
	calDay := 0
	calMonth := 1
	timeThisTimeStart, timeThisTimeEnd = utils.GetTimeUTCThisMonth(-480)

	timeLastTimeStart = timeThisTimeStart.AddDate(0, -calMonth, -calDay)
	timeLastTimeEnd = timeLastTimeStart.AddDate(0, calMonth, calDay)
	t.Log(timeThisTimeStart, timeThisTimeEnd)
	t.Log(timeLastTimeStart, timeLastTimeEnd)
}

func TestGetTimeUTCThisWeek(t *testing.T) {
	// utils.GetTimeUTCToday(-480)
	var timeThisTimeStart, timeThisTimeEnd time.Time
	var timeLastTimeStart, timeLastTimeEnd time.Time
	calDay := 7
	calMonth := 0
	timeThisTimeStart, timeThisTimeEnd = utils.GetTimeUTCThisWeek(-480)

	timeLastTimeStart = timeThisTimeStart.AddDate(0, -calMonth, -calDay)
	timeLastTimeEnd = timeLastTimeStart.AddDate(0, calMonth, calDay)
	t.Log(timeThisTimeStart, timeThisTimeEnd)
	t.Log(timeLastTimeStart, timeLastTimeEnd)
}
