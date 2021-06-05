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
	Name         string
	FindFunc     echo.HandlerFunc
	GetFunc      echo.HandlerFunc
	CreateFunc   echo.HandlerFunc
	App          *App
	ContextValue map[string]interface{}
}

func NewService(name string, find echo.HandlerFunc, get echo.HandlerFunc, create echo.HandlerFunc) *Service {
	return &Service{
		Name:         name,
		FindFunc:     find,
		GetFunc:      get,
		CreateFunc:   create,
		ContextValue: make(map[string]interface{}),
	}
}

func (svc *Service) AddContextValue(key string, value interface{}) error {
	svc.ContextValue[key] = value
	return nil
}

func (svc *Service) Find(c echo.Context) error {
	if svc.FindFunc != nil {
		for key, value := range svc.ContextValue {
			c.Set(key, value)
		}
		svc.FindFunc(c)
	} else {
		return errors.New("no method find")
	}
	return nil
}

func (svc *Service) Get(c echo.Context) error {
	if svc.GetFunc != nil {
		for key, value := range svc.ContextValue {
			c.Set(key, value)
		}
		svc.GetFunc(c)
	} else {
		return errors.New("no method find")
	}
	return nil
}

func (svc *Service) Create(c echo.Context) error {
	if svc.CreateFunc != nil {
		for key, value := range svc.ContextValue {
			c.Set(key, value)
		}
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
