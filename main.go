package main

import (
	"fmt"
	"gin-study/setting"
)

func main() {
	setting.Init()
	fmt.Printf("%+v\n", setting.Config)
	fmt.Printf("%+v\n", setting.Config.LogConfig)
	fmt.Printf("%+v\n", setting.Config.MysqlConfig)
	fmt.Printf("%+v\n", setting.Config.RedisConfig)
}
