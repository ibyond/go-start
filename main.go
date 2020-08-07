package main

import (
	"github.com/ibyond/go-start/core"
	"github.com/ibyond/go-start/global"
	"github.com/ibyond/go-start/initialize"
)

func main() {
	switch global.GstConfig.System.DbType {
	case "mysql":
		initialize.Mysql()
	//case "sqlite":
	//	initialize.Sqlite()
	default:
		initialize.Mysql()
	}
	sqlDb, _ := global.GstDb.DB()
	defer sqlDb.Close()

	core.RunServer()

}
