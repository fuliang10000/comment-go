package config

import "commentService/pkg/gormx"

type Config struct {
	Mysql      gormx.Config `json:"mysql" config:"mysql"`
	ServerPort int          `json:"server_port" config:"server_port,default:9001"`
}
