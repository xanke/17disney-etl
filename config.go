package main

import (
	"os"
)

func getString(name string, def interface{}) string {
	val := os.Getenv(name)
	if val == "" && def == nil {
		panic(name)
	}

	if val != "" {
		return val
	}
	return def.(string)
}

// 服务器配置
var (
	mysqlUrl         = getString("MYSQL_URL", "http://127.0.0.1")
	redisUrl         = getString("REDIS_URL", "http://127.0.0.1")
)
