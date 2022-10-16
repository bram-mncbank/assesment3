package router

import (
	"assesment3/webserver/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		controllers.Render(ctx, "./webserver/views/index.html", map[string]interface{}{
			"Title": "Assesment 3",
		})
	})
	router.GET("/weather", controllers.GetWeathers)

	return router
}
