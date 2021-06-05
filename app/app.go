package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Installable interface {
	Install(*App) error
}

type App struct {
	Echo     *echo.Echo
	Config   Config
	Services map[string]*Service
	Models   map[string]*Model
	Db       *gorm.DB
}

func NewApp(Echo *echo.Echo) *App {
	return &App{Echo: Echo, Config: Config{}, Services: make(map[string]*Service)}
}

func (app *App) RegisterService(svc *Service) error {
	app.Services[svc.Name] = svc
	// register groups
	groups := app.Echo.Group(svc.Name)
	// register Get
	if svc.FindFunc != nil {
		groups.GET("", svc.Find)
	}
	if svc.GetFunc != nil {
		groups.GET("/:id", svc.Get)
	}
	if svc.CreateFunc != nil {
		groups.POST("", svc.Create)
	}
	return nil
}

func (app *App) RegisterModel(model Model) error {
	if err := app.Db.AutoMigrate(model); err != nil {
		return err
	}
	return nil
}

func (app *App) RegisterDatabase(db *gorm.DB) error {
	app.Db = db
	return nil
}

func (app *App) Start() error {
	return app.Echo.Start(fmt.Sprintf(":%s", app.Config.Port))
}

func (app *App) Use(module Installable) error {
	return module.Install(app)
}
