package main

import (
	"b30northwindapi/config"
	"b30northwindapi/server"
	"log"
	"os"
)

func main() {
	log.Println("Starting Northwind API")

	log.Println("Initialiazing Configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)

	//log.Println(dbHandler)

	log.Println("Initializig HTTP sever")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if env != "" {
		return "northwind-" + env
	}

	return "northwind"
}
