package main

import (
	"flag"
	"prodp/api"
	"prodp/data"
)

var (
	isProduction = flag.Bool("production", false, "Indicates production environment")
	migrateData  = flag.Bool("data", false, "Setup tables, enable on first run")
)

func main() {
	flag.Parse()

	dataRepo, err := data.NewClient()
	if err != nil {
		panic(err)
	}

	if *migrateData {
		if err := dataRepo.Migrate(); err != nil {
			panic(err)
		}
	}

	server := api.Server{Repository: dataRepo}

	server.Serve(api.Options{Production: *isProduction})
}
