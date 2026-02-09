package model

type AgentList struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Md5Key string `json:"md5_key"`
	AesKey string `json:"aes_key"`
}
