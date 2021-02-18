package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iwara/database"
	"iwara/models"
	"net/http"
	"strconv"
)

type VideoController struct{}

func (v *VideoController) Get(c *gin.Context) {
	pg := c.Query("page")
	page, _ := strconv.Atoi(pg)
	if page--; page < 0 {
		page = 0
	}
	var vs models.Videos
	database.Sql(func(db *gorm.DB) {
		db.Limit(models.PerPage).Offset(models.PerPage * page).Find(&vs)
	})
	c.JSON(http.StatusOK, vs)
}

func (v *VideoController) Show(c *gin.Context) {
	id := c.Param("video")
	var d models.Video
	database.Sql(func(db *gorm.DB) {
		db.First(&d, id)
	})
	c.JSON(http.StatusOK, d)
}
