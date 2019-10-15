package load

import (
	"fmt"

	"github.com/Oleg-Skalozub/testtask/src/infrastructure/config"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/db"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/logger"

	"github.com/pkg/errors"
)

// LoaderList is a collection of Load() functions
type LoaderList []struct {
	name string
	load func() error
}

// BasicLoaders contains load() functions of internal components.
// Order of loaders matters!
var (
	Loaders = LoaderList{
		{"config", config.Load},
		{"logger", logger.Load},
		{"sqlbd", db.Load},
	}
	Unloaders = LoaderList{
		{"sqlbd", db.UnLoad},
	}
)

// LoadApplicationServices ...
func LoadApplicationServices() error {
	return executeLoaders(Loaders)
}

// UnloadApplicationServices ...
func UnloadApplicationServices() error {
	return executeLoaders(Unloaders)
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
