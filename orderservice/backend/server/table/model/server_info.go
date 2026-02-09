package model

type ServerInfo struct {
	Code           string    `json:"code" db:"code"`
	Ip             string    `json:"ip" db:"ip"`
	Addresses      Addresses `json:"addresses" db:"addresses"`
	AddressesBytes []byte    `json:"-"`
	IsEnabled      bool      `json:"is_enabled" db:"is_enabled"`
}

type Addresses struct {
	Notification string `json:"notification" db:"notification"`
}
