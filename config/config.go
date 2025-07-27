package config

type Config struct {
	EndPoint string `json:"endpoint"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	
}