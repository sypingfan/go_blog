package main

import (
	"admin/core"
	"admin/global"
	"fmt"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
