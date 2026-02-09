package module_lib_test

import (
	"monitorservice/cmd/baseserver/api_module/module_lib"
	"testing"
)

func TestGenerateTimeStrWeek(t *testing.T) {
	t.Log(module_lib.GenerateTimeStrWeek())
}

func TestGenerateTableNameMonth(t *testing.T) {
	t.Log(module_lib.GenerateMonthPeriodOfTime("15min", -480))
}

func TestGenerateWeekPeriodOfTime(t *testing.T) {
	t.Log(module_lib.GenerateWeekPeriodOfTime("15min", -480))

}
