package core

import (
	"admin/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置Mysql，取消gorm链接")
		return nil
	}
	dsn := global.Config.Mysql.Dns()
	var mysqllogger logger.Interface
	if global.Config.System.Env == "debug" {
		//开发环境显示所有的SQL
		mysqllogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqllogger = logger.Default.LogMode(logger.Error) //只打印错误的Sql
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqllogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql链接失败", dsn))
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDb.SetMaxOpenConns(100)              //最多可容纳
	sqlDb.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
