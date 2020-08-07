package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ibyond/go-start/global"
	"strings"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-CSRF-Token", "AccessToken", "Authorization", "Token"}
	// 运行再Release模式下才会进行跨域保护 保证开发过程中不会被跨域困扰~
	origin := global.GstConfig.System.Origins
	origins := strings.Split(origin, ",")
	if gin.Mode() == gin.ReleaseMode {
		config.AllowOrigins = origins
	} else {
		config.AllowAllOrigins = true
	}
	config.AllowCredentials = true
	return cors.New(config)
}
