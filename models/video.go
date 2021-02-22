package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Video struct {
	Model
	Url    string `json:"url"`
	Poster string `json:"poster"`
	Title  string `json:"title"`
	View   int    `json:"view"`
	Star   int    `json:"star"`
	HashId string `json:"hash_id"`
}

type Videos []Video

func When(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		flag := false
		star := c.Query("star")
		switch {
		case star == "asc":
			fallthrough
		case star == "desc":
			db.Order("star " + star)
		default:
			flag = true
		}
		log.Println(star)
		view := c.Query("view")
		switch {
		case view == "asc":
			fallthrough
		case view == "desc":
			db.Order("view " + view)
		default:
			flag = true
		}
		if flag {
			db.Order("created_at desc")
		}
		title := c.Query("title")
		if len(title) > 0 {
			db.Where("title LIKE ?", "%"+title+"%")
		}

		return db
	}
}
