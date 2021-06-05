package env

import (
	"log"

	"github.com/arifusr/online_store/app"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Env struct {
}

func (e *Env) Install(a *app.App) error {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := envdecode.Decode(&a.Config); err != nil {
		log.Fatal(err)
	}

	return nil
}
