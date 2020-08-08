package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	resp "github.com/ibyond/go-start/app/model"
	"github.com/ibyond/go-start/global"
	"github.com/ibyond/go-start/global/response"
	"github.com/mojocn/base64Captcha"
)

type CaptchaController struct{}

var store = base64Captcha.DefaultMemStore

func (p CaptchaController) Captcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GstConfig.Captcha.ImgHeight, global.GstConfig.Captcha.ImgWidth, global.GstConfig.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkDetailed(resp.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
