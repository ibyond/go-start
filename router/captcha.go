package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibyond/go-start/app/controllers"
)

func InitCaptchaRouter(Router *gin.RouterGroup) {
	CaptchaRoute := Router.Group("")
	{
		captCtl := new(controllers.CaptchaController)
		CaptchaRoute.POST("captcha", captCtl.Captcha)
	}
}
