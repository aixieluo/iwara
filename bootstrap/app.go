package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12/sessions"
	"log"
	"time"
)

type Configurator func(app *App)

type App struct {
	*gin.Engine
	AppName string
	AppSpawnDate time.Time

	Session *sessions.Sessions
	Config int
}

func New(appName string, cgs ...Configurator) *App {
	// 切换正式版
	// gin.SetMode(gin.ReleaseMode)

	t:=time.Now()
	app := &App{
		Engine: gin.Default(),
		AppName:appName,
		AppSpawnDate: t,
	}
	log.Println(t.String())

	app.Configure(cgs...)

	return app
}

func (app *App) Configure(cgs ...Configurator)  {
	for _, cg := range cgs {
		cg(app)
	}
}

func (app *App) Bootstrap() *App {
	app.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(10 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	app.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(10 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	app.GET("/long_sync2", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(10 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	return app
}
