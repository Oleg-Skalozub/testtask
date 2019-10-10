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
		config.Config.DBConfig.DBHost, config.Config.DBConfig.DBPort, config.Config.DBConfig.DBUser, config.Config.DBConfig.DBName, config.Config.DBConfig.DBPassword)

	db, err := gorm.Open(config.Config.DBConfig.DBDialect, conf)
	if err != nil {
		return nil
	}

	db.DB().SetMaxOpenConns(config.Config.DBConfig.DBConn)
	db.DB().SetMaxIdleConns(config.Config.DBConfig.DBIdleConn)

	SQLBD = sqlDB{db: db}

	return err
}

// UnLoad ...
func UnLoad() error {
	return SQLBD.Close()
}
