package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibyond/go-start/app/controllers"
)

func InitCaptchaRouter(Router *gin.RouterGroup) {
	BaseRoute := Router.Group("")
	{
		captCtl := new(controllers.CaptchaController)
		BaseRoute.POST("captcha", captCtl.Captcha)
	}
}
