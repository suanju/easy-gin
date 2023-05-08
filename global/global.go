package global

import (
	log "easy-gin/global/logrus"
	"github.com/sirupsen/logrus"
)

func init() {
	Logger = log.ReturnsInstance()

}

var (
	Logger *logrus.Logger
)
