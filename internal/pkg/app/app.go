package app

import (
	"fmt"
	"github.com/cucumberjaye/qtsoft/configs"
	"github.com/cucumberjaye/qtsoft/internal/app/handler"
	"github.com/cucumberjaye/qtsoft/internal/app/service"
	"net/http"
)

type App struct {
	h *handler.Handler
	s *service.Service
}

func New() *App {
	app := &App{}

	app.s = service.New()

	app.h = handler.New(app.s)

	return app
}

func (a *App) Run() error {
	fmt.Println("server running")

	host := configs.Domain + ":" + configs.Port
	return http.ListenAndServe(host, a.h.InitRoutes())
}
