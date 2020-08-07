package core

import (
	"github.com/ibyond/go-start/global"
	"github.com/ibyond/go-start/initialize"
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
}
