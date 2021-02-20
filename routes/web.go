package routes

import (
	"github.com/gin-gonic/gin"
	"iwara/bootstrap"
	"iwara/http/controllers"
	"iwara/untils/spider"
	"net/http"
)

func Configure(app *bootstrap.App) {
	app.GET("total", func(c *gin.Context) {
		s := new(spider.Spider)
		s.Start()
		c.JSON(http.StatusOK, "ok")
	})
	app.GET("/video", (&controllers.VideoController{}).Get)
	app.GET("/video/:video/show", (&controllers.VideoController{}).Show)
	app.GET("test", func(c *gin.Context) {
	})
}
