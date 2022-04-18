package main

import (
	"flag"
	"fmt"
	"log"
	"students/config"
	"students/db"
	"students/routers"

	"github.com/gin-gonic/gin"
)

var (
	confPath = flag.String("conf", "../conf/config.yaml", "config file path")
)

func main() {
	flag.Parse()

	err := config.Init(*confPath)
	if err != nil {
		panic(err)
	}

	err = db.InitMysql(config.Conf)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routers.InitRouter(r)

	log.Fatalln(r.Run(fmt.Sprintf(":%d", config.Conf.Web.Port)))
}
