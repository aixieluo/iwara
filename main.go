package main

import (
	"iwara/bootstrap"
	"iwara/routes"
	"iwara/untils"
	"log"
	"time"
)

func newApp() *bootstrap.App {
	app := bootstrap.New("Iwara")
	app.Configure(routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Bootstrap()
	untils.Schedules(func() {
		log.Println(time.Now().String())
	})

	err := app.Run(":3939")
	if err != nil {
		log.Println(err)
	}
}
