package database

import (
	"github.com/arifusr/online_store/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseSqlite struct {
}

func (d *DatabaseSqlite) Install(a *app.App) error {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	a.RegisterDatabase(db)

	return nil
}
