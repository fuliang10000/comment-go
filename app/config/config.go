package config

import (
	"commentService/pkg/configx"
	"commentService/pkg/gormx"
	"commentService/rpc"
	"flag"
)

type Config struct {
	ServerPort int          `json:"server_port" config:"server_port,default:9001"`
	Mysql      gormx.Config `json:"mysql" config:"mysql"`
	Rpc        rpc.Config   `json:"rpc" config:"rpc"`
}

var f = flag.String("c", "./etc/config.yaml", "config path")

// NewConfig 加载配置文件
func NewConfig() *Config {
	var c Config
	flag.Parse()
	configx.MustLoad(*f, &c)
	return &c
}

func RpcConfig(c *Config) *rpc.Config {
	return &c.Rpc
}
