package core

import (
	"fmt"
	"github.com/ibyond/go-start/global"
	"github.com/ibyond/go-start/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GstConfig.System.UseMultipoint {
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GstConfig.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GstLog.Debug("server run success on ", address)

	fmt.Println("欢迎使用 GO-START", address)
	global.GstLog.Error(s.ListenAndServe())
}
