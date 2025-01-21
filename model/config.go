package model

type DbConfig struct {
	User     string `json:"user"`
	PassWord string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Schema   string `json:"schema"`
}
