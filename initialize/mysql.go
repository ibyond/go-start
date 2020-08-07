package initialize

import (
	"github.com/ibyond/go-start/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Mysql() {
	var db *gorm.DB
	var err error

	admin := global.GstConfig.Mysql
	dsn := admin.Username + ":" + admin.Password + "@tcp(" + admin.Path + ")/" + admin.Dbname + "?" + admin.Config

	if admin.LogMode {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,   // 慢 SQL 阈值
				LogLevel:      logger.Silent, // Log level
				Colorful:      false,         // 禁用彩色打印
			},
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
	} else {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		global.GstLog.Error("Mysql 启动异常 ", err)
	} else {
		global.GstDb = db
		if sqlDb, err := global.GstDb.DB(); err != nil {
			global.GstLog.Error("Mysql 启动异常: ", err)
		} else {
			sqlDb.SetMaxIdleConns(admin.MaxIdleConns)
			sqlDb.SetMaxOpenConns(admin.MaxOpenConns)
		}
	}

}
