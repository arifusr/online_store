package main

import (
	"github.com/arifusr/online_store/app"
	"github.com/arifusr/online_store/modules"
	"github.com/labstack/echo/v4"
)

func main() {

	// init app object
	E := echo.New()
	App := app.NewApp(E)

	// install modules
	M := &modules.Modules{}
	App.Use(M)

	App.Start()

}
