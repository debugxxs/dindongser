package main

import (
	"dingdongser/router"
	"dingdongser/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg,err:=tools.ParsingConfig("./config/config.json")
	if err !=nil{
		panic(err.Error())
	}
	tools.InitDbEngine(cfg)
	app := gin.Default()
	router.LoadRouter(app)
	_ = app.Run(cfg.AppHost + ":" + cfg.AppPort)
}
