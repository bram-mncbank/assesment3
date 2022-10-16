package webserver

import (
	"assesment3/webserver/configs"
	"assesment3/webserver/router"
)

func Start() {
	// engine := gin.Default()
	router.CreateRouter().Run(configs.PORT)
}
