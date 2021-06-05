package app

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type ServiceInterface interface {
	Find(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Patch(c echo.Context) error
	Remove(c echo.Context) error
}

type Service struct {
	Name       string
	FindFunc   echo.HandlerFunc
	GetFunc    echo.HandlerFunc
	CreateFunc echo.HandlerFunc
	App        *App
}

func NewService(name string, a *App, find echo.HandlerFunc, get echo.HandlerFunc, create echo.HandlerFunc) *Service {
	return &Service{
		Name:       name,
		FindFunc:   find,
		GetFunc:    get,
		CreateFunc: create,
		App:        a,
	}
}

func (svc *Service) Find(c echo.Context) error {
	if svc.FindFunc != nil {
		svc.FindFunc(c)
	} else {
		return errors.New("no method find")
	}
	return nil
}

func (svc *Service) Get(c echo.Context) error {
	if svc.GetFunc != nil {
		svc.GetFunc(c)
	} else {
		return errors.New("no method find")
	}
	return nil
}

func (svc *Service) Create(c echo.Context) error {
	if svc.CreateFunc != nil {
		c.Set("App", svc.App)
		svc.CreateFunc(c)
	} else {
		return errors.New("no method find")
	}
	return nil
}

func (svc *Service) Update(c echo.Context) error {
	return nil
}

func (svc *Service) Patch(c echo.Context) error {
	return nil
}

func (svc *Service) Remove(c echo.Context) error {
	return nil
}
