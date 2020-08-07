package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/ibyond/go-start/config"
	"gorm.io/gorm"

	logging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	GstDb     *gorm.DB
	GstRedis  *redis.Client
	GstConfig config.Server
	GstVp     *viper.Viper
	GstLog    *logging.Logger
)
