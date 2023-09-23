package main

import (
	"admin/core"
	"admin/global"
	"fmt"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnf("哈哈哈")
	global.Log.Error("哈哈哈")
	global.Log.Infof("哈哈哈")
	global.Log.Debugf("哈哈哈")
	// 连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
