package main

import (
	"admin/core"
	"admin/flag"
	"admin/global"
	"admin/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	//global.Log.Warnf("哈哈哈")
	//global.Log.Error("哈哈哈")
	//global.Log.Infof("哈哈哈")
	//global.Log.Debugf("哈哈哈")
	// 连接数据库
	global.DB = core.InitGorm()

	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()

	addr := global.Config.System.Adr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

}
