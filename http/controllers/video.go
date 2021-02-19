package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iwara/database"
	"iwara/models"
	"iwara/untils/spider"
	"net/http"
)

type VideoController struct{}

func (v *VideoController) Get(c *gin.Context) {
	var vs models.Videos
	database.Sql(func(db *gorm.DB) {
		db.Scopes(models.Paginate(c)).Find(&vs)
	})
	c.JSON(http.StatusOK, vs)
}

func (v *VideoController) Show(c *gin.Context) {
	id := c.Param("video")
	var d models.Video
	database.Sql(func(db *gorm.DB) {
		db.First(&d, id)
	})
	url := spider.Video(d.Url)
	c.HTML(http.StatusOK, "show.tpl", gin.H{
		url: url,
	})
}
