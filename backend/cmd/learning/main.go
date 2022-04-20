package main

import (
	"flag"
	"fmt"
	"learning/config"
	"learning/db/cache"
	"learning/db/mysql"
	"learning/routers"
	"log"

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

	r := gin.Default()

	routers.InitRouter(r)

	log.Fatalln(r.Run(fmt.Sprintf(":%d", config.Conf.Web.Port)))
}
