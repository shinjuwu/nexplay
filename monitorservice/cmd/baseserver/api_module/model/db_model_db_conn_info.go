package model

import (
	"fmt"
	"time"
)

const (
	// drivername
	DBType_PostgreSQL = "postgres"
	// DBType_MySQL      = "mysql"
)

// // database 基本連線設定參數
// type DatabaseConfig struct {
// 	DriverName string `json:"driver_name"`
// 	Host       string `json:"host"`
// 	Port       int    `json:"port"`
// 	Username   string `json:"username"`
// 	Password   string `json:"password"`
// 	Database   string `json:"database"`
// }

//*****************************************************************************

type PostgreSQLConnInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
	Driver   string `json:"driver"`
}

func (c *PostgreSQLConnInfo) GetConnInfoString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, c.Database, c.SSLMode)
}

//*****************************************************************************

type DBConnInfo struct {
	ID             string             `json:"id"`          // p key
	AddressesBytes []byte             `json:"-"`           // used to parse
	Addresses      PostgreSQLConnInfo `json:"addresses"`   // json, db conn info
	Info           string             `json:"info"`        // 備註
	IsEnabled      bool               `json:"is_enabled"`  // 是否啟用
	CreateTime     time.Time          `json:"create_time"` // 創建時間
	UpdateTime     time.Time          `json:"update_time"`
	DisableTime    time.Time          `json:"disable_time"`
}
