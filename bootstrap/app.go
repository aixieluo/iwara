package bootstrap

import (
	"github.com/Masterminds/sprig"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12/sessions"
	"html/template"
	"iwara/untils"
	"log"
	"time"
)

type Configurator func(app *App)

type App struct {
	*gin.Engine
	AppName      string
	AppSpawnDate time.Time

	Session *sessions.Sessions
	Config  int
}

func New(appName string, cgs ...Configurator) *App {
	// 切换正式版
	// gin.SetMode(gin.ReleaseMode)

	t := time.Now()
	app := &App{
		Engine:       engine(),
		AppName:      appName,
		AppSpawnDate: t,
	}

	app.Configure(cgs...)

	return app
}

func engine() *gin.Engine {
	e := gin.Default()
	tpl := template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseGlob("views/*/*.html"))
	e.SetHTMLTemplate(tpl)
	return e
}

func (app *App) Configure(cgs ...Configurator) {
	for _, cg := range cgs {
		cg(app)
	}
}

func (app *App) Bootstrap() *App {
	untils.Schedule(func() {
		log.Println(time.Now().String())
	})
	return app
}
