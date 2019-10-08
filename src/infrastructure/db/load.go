package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/Oleg-Skalozub/testtask/src/infrastructure/config"
)

// SQLBD main object for working with DB
var SQLBD SQLBDer

// Load ...
func Load() error {

	conf := fmt.Sprintf("sslmode=disable host=%s port=%d user=%s dbname=%s password=%s",
		config.Config.DBHost, config.Config.DBPort, config.Config.DBUser, config.Config.DBName, config.Config.DBPassword)

	db, err := gorm.Open(config.Config.DBDialect, conf)

	SQLBD = sqlDB{db: db}
	return err
}

// UnLoad ...
func UnLoad() error {
	return SQLBD.Close()
}
