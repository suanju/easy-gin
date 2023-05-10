package global

import (
	"easy-gin/global/config"
	"easy-gin/global/database/mysql"
	RedisDbFun "easy-gin/global/database/redis"
	log "easy-gin/global/logrus"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	Logger = log.ReturnsInstance()
	RedisDb = RedisDbFun.ReturnsInstance()
	Db = mysql.ReturnsInstance()
	Config = config.ReturnsInstance()

}

var (
	Logger  *logrus.Logger
	Config  *config.Info
	Db      *gorm.DB
	RedisDb *redis.Client
)
