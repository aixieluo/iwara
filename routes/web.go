package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iwara/bootstrap"
	"iwara/database"
	"iwara/http/controllers"
	"iwara/models"
	"iwara/untils/spider"
	"log"
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
		v := &models.Video{}

		database.Sql(func(db *gorm.DB) {
			db.Limit(1).Find(&v)
		})
		log.Println(v.ID)
	})
}
