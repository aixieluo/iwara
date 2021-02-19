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
	app.LoadHTMLGlob("templates/*/*")
	return app
}
