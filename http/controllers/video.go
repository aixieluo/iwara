package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
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
	var video models.Video
	database.Sql(func(db *gorm.DB) {
		db.First(&video, id)
	})
	sc :=spider.NewCollector()
	var body interface{}
	sc.OnResponse(func(res *colly.Response) {
		_ = json.Unmarshal(res.Body, &body)
	})
	_ = sc.Visit("https://ecchi.iwara.tv/api/video/" + video.HashId)
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"sources": body,
		"name":    "video",
	})
}
