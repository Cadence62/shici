package utils

import (
	"github.com/spf13/viper"
	"shici/config"
)

var INITCONFIG *config.Config

// 加载配置
func InitConfig() (c *config.Config, err error) {
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")
	v.SetConfigType("yaml")
	if err = v.ReadInConfig(); err != nil {
		return
	}
	c = &config.Config{}
	c.Rabbitmq.Host = v.GetString("rabbitmq.host")
	c.Rabbitmq.PingExchange = v.GetString("rabbitmq.pingexchange")
	return
}
