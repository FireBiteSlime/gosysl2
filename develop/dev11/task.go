package main

import (
	"develop/dev11/task11/calendar"
	"develop/dev11/task11/routes"
	"develop/dev11/task11/server"
	"log"
)

func main() {
	calendar := calendar.New()
	routes.IntitRoutes(calendar)
	server := new(server.Server)
	err := server.Run("8080", 10)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Server start...")
}
