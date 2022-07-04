package main

import (
	"log"
	"sandbox/database"
	"sandbox/service"
	"sandbox/web"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to db %v", err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Unable to migrate db %v", err)

	}

	c := service.New(db)
	service.SeedData(c)
	web.RegisterAPIRoutes(c)
	c.Web.Run()
}
