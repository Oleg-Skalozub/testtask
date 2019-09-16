package main

import (
	"flag"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/config"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/load"
	"github.com/Oleg-Skalozub/testtask/src/router"
	"log"
	"net/http"
	"os"
)

var configFile *string
var logFatal = log.New(os.Stderr, "ERROR:\n", 0).Fatalf

func init() {
	configFile = flag.String("config", "", "Configuration file in JSON-format")
}

func main() {
	flag.Parse()

	if len(*configFile) > 0 {
		config.FilePath = *configFile
	}

	err := load.LoadApplicationServices()
	if err != nil {
		logFatal("Failed to initialize : %v", err)
	}

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
