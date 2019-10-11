package logger

import (
	"log"
	"os"

	"github.com/Oleg-Skalozub/testtask/src/infrastructure/config"
)

// Log ...
var Log Logger

// Load ...
func Load() error {

	file, err := os.OpenFile(config.Config.LogConfig.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open log file")
	}

	log.SetOutput(file)
	logg := log.New(file, config.Config.LogConfig.Prefix, log.Lshortfile)
	logg.SetFlags(1)
	Log = &logger{logMain: logg}
	return nil
}
