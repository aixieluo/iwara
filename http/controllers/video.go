package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"gorm.io/gorm"
	"iwara/database"
	"iwara/http/resource"
	"iwara/models"
	"iwara/untils/spider"
	"net/http"
)

type VideoController struct{}

func (v *VideoController) Get(c *gin.Context) {
	var vs models.Videos
	var total int64
	database.Sql(func(db *gorm.DB) {
		db.Scopes(models.Paginate(c), models.When(c)).Find(&vs)
		db.Model(&models.Video{}).Count(&total)
	})
	c.JSON(http.StatusOK, resource.Factory(vs).SetMeta(gin.H{
		"total": total,
	}))
}

func (v *VideoController) Show(c *gin.Context) {
	id := c.Param("video")
	var video models.Video
	database.Sql(func(db *gorm.DB) {
		db.First(&video, id)
	})
	if len(video.HashId) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "无此视频",
		})
	}
	sc := spider.NewCollector()
	var body interface{}
	sc.OnResponse(func(res *colly.Response) {
		_ = json.Unmarshal(res.Body, &body)
	})
	_ = sc.Visit("https://ecchi.iwara.tv/api/video/" + video.HashId)
	c.JSON(http.StatusOK, resource.Factory(body))
}
