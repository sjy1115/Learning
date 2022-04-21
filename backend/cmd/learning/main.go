package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"learning/config"
	"learning/db/cache"
	"learning/db/mysql"
	"learning/middleware"
	"learning/pkg/captcha"
	"learning/pkg/log"
	"learning/routers"

	"github.com/gin-gonic/gin"
)

var (
	confPath = flag.String("f", "", "config file path")
)

func main() {
	flag.Parse()

	err := config.Init(*confPath)
	if err != nil {
		panic(err)
	}

	err = mysql.InitMysql(config.Conf)
	if err != nil {
		panic(err)
	}
	cache.InitRedis(&config.Conf.Redis)

	initLog(config.Conf)
	captcha.InitCaptcha()
	middleware.InitIgnoreUrl(config.Conf.IgnoreUrl)

	r := gin.Default()

	routers.InitRouter(r)

	logrus.Fatal(r.Run(fmt.Sprintf(":%d", config.Conf.Web.Port)))
}

func initLog(cfg *config.Config) {
	logPath := cfg.Log.LogPath
	if cfg.Web.Docker {
		logPath = cfg.Log.DockerLogPath
	}

	log.InitLog(cfg.Log.Level, logPath)
}
