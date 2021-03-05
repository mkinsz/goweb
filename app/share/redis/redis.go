package redis

import "time"

type Redis struct {
	Host        string        `json:"Host"`
	Password    string        `json:"Password"`
	MaxIdle     int           `json:"MaxIdle"`
	MaxActive   int           `json:"MaxActive"`
	IdleTimeout time.Duration `json:"IdleTimeout"`
}
