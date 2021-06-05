package database

import (
	"fmt"

	"github.com/arifusr/online_store/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabasePostgres struct {
}

func (d *DatabasePostgres) Install(a *app.App) error {

	sqlCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		a.Config.ConfigDatabase.Host,
		a.Config.ConfigDatabase.Port,
		a.Config.ConfigDatabase.Username,
		a.Config.ConfigDatabase.Password,
		a.Config.ConfigDatabase.Name,
	)

	db, err := gorm.Open(postgres.Open(sqlCfg), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic("failed to connect database")
	}

	a.RegisterDatabase(db)

	return nil
}
