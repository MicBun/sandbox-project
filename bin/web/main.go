package main

import (
	"log"
	"math/rand"
	"sandbox/database"
	"sandbox/docs"
	"sandbox/service"
	"sandbox/web"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to db %v", err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Unable to migrate db %v", err)

	}

	contact := "Checkout my Github: https://github.com/MicBun\n\n" +
		"Checkout my Linkedin: https://www.linkedin.com/in/MicBun\n\n"

	description := "This is a sample API server.\n\n" +
		"Use this server to learn about API design and implementation.\n\n" +
		contact

	docs.SwaggerInfo.Title = "Sandbox API"
	docs.SwaggerInfo.Description = description

	c := service.New(db)
	service.SeedData(c)
	web.RegisterAPIRoutes(c)
	c.Web.Run()
}