package logger

import (
	"log"
	"os"
)

// Log ...
var Log Logger

// Load ...
func Load() error {

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed")
	}

	log.SetOutput(file)
	logg := log.New(file, "test_task", log.Lshortfile)
	logg.SetFlags(1)
	Log = &logger{logMain: logg}
	return nil
}
