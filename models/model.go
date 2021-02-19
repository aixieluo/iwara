package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey;column:id" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Param("page"))
		if page < 1 {
			page = 1
		}

		perPage, _ := strconv.Atoi(c.Param("perPage"))
		switch {
		case perPage > 100:
			perPage = 100
		case perPage < 1:
			perPage = 10
		}
		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
