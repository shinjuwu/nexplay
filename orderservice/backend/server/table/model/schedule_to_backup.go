package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type ScheduleToBackup struct {
	Id               string    `json:"id"`
	DataKeepingTable string    `json:"data_keeping_table"`
	DataKeepingDay   int       `json:"data_keeping_day"`
	IsEnabled        bool      `json:"is_enabled"`
	CreateTime       time.Time `json:"create_time"`
	UpdateTime       time.Time `json:"update_time"`
	DisableTime      time.Time `json:"disable_time"`
	LastExecTime     time.Time `json:"last_exec_time"`
}

func NewEmptyScheduleToBackup() *ScheduleToBackup {

	genUUID, _ := uuid.NewV4()

	return &ScheduleToBackup{
		Id:               genUUID.String(),
		DataKeepingTable: "",
		DataKeepingDay:   -1,
		IsEnabled:        false,
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
		DisableTime:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		LastExecTime:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func NewScheduleToBackup(dataKeepingTable string, dataKeepingDay int, dataKeepingMinOffset int, isEnabled bool) *ScheduleToBackup {

	genUUID, _ := uuid.NewV4()

	return &ScheduleToBackup{
		Id:               genUUID.String(),
		DataKeepingTable: dataKeepingTable,
		DataKeepingDay:   dataKeepingDay,
		IsEnabled:        isEnabled,
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
		DisableTime:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		LastExecTime:     time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}
