package main

import (
	"log"
	"os"
	"github.com/timurgulov/calc_go/api"
)

func main() {
	app := api.New()

	if len(os.Args) > 1 && os.Args[1] == "server" {
		if err := app.RunServer(); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	} else {
		if err := app.RunCLI(); err != nil {
			log.Fatal("Failed to run CLI:", err)
		}
	}
}