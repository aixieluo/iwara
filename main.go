package main

import (
	"iwara/bootstrap"
	"iwara/routes"
	"log"
)


func Default() *bootstrap.App {
	app := bootstrap.New("Iwara")
	app.Configure(routes.Configure)
	app.Bootstrap()
	return app
}

func main() {
	app := Default()
	err := app.Run(":3939")
	if err != nil {
		log.Println(err)
	}
}
