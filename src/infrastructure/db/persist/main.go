package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Oleg-Skalozub/testtask/src/domain/entity"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/config"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/db"

	"github.com/pkg/errors"
)

var configFile *string

func init() {
	configFile = flag.String("config", "config.json", "Configuration file in JSON-format")
}

// LoaderList is a collection of Load() functions
type LoaderList []struct {
	name string
	load func() error
}

func main() {
	flag.Parse()

	if len(*configFile) > 0 {
		config.FilePath = *configFile
	}

	var BasicLoaders = LoaderList{
		{"config", config.Load},
		{"sqlbd", db.Load},
	}

	err := executeLoaders(BasicLoaders)

	if err != nil {
		log.Fatal(err)
	}

	start()
}

func executeLoaders(loaders LoaderList) error {
	for _, loader := range loaders {
		err := loader.load()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to execute %s.Load()", loader.name))
		}
	}
	return nil
}

func start() {
	db.SQLBD.AutoMigrate(&entity.DataDB{})
}
