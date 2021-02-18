 package models

 import "gorm.io/gorm"

 const PerPage  = 15

 type Model struct {
 	*gorm.Model
 }
