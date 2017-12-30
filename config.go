package main

// 服务器配置
var (
	mysqlUrl         = getString("MYSQL_URL", nil),
	redisUrl         = getString("REDIS_URL", nil)
)
