package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ibyond/go-start/global"
	"github.com/ibyond/go-start/middleware"
	"github.com/ibyond/go-start/router"
)

func Routers() *gin.Engine {
	if global.GstConfig.System.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	var Router = gin.Default()
	global.GstLog.Debug("use middleware logger")
	Router.Use(middleware.Cors())
	//Router.Use(middleware.LoadTls())

	ApiGroup := Router.Group("")
	router.InitCaptchaRouter(ApiGroup)
	return Router
}
