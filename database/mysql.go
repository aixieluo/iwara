package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var DbPool chan *gorm.DB
var poolMax int = 8

type sqlFunc func(db *gorm.DB)

func init() {
	DbPool = make(chan *gorm.DB, poolMax)
	for i := 0; i < poolMax; i++ {
		DbPool <- connect()
	}
}

func pop() *gorm.DB {
	return <-DbPool
}

func push(db *gorm.DB)  {
	DbPool<-db
}

func connect() *gorm.DB {
	dsn := "root:123123@tcp(database)/iwara?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic("db can't connect!")
	}
	return db
}

func Sql(fn sqlFunc) {
	db := pop()
	fn(db)
	push(db)
}
