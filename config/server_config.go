package config

const (
	//sql server config
	PostgreIP       = "localhost"
	PostgrePort     = "5432"
	PostgreUser     = "postgres"
	PostgreDbname   = "shici"
	PostgreSSL      = "disable"
	PostgrePassword = "123qwe"

	// redis server config
	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = "123qwe"
	RedisDB       = 0

	// redis key name
	RedisKeyTodayOne = "todayone"
)

type Config struct {
	Rabbitmq Rabbitmq
}

// rabbitmq配置
type Rabbitmq struct {
	Host         string
	PingExchange string
}
