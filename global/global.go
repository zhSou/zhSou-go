package global

import (
	"github.com/zhSou/zhSou-go/core/config"
)

var Config *config.Config

func InitGlobal(conf *config.Config) {
	Config = conf
}
