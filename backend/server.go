package main

import (
	"github.com/Qwiri/gobby/pkg/gobby"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()
	g := gobby.New(app)
	log.Println(g)

	initQuizRoutes(app)
	initUserRoutes(app)

	if err := app.Listen(":8081"); err != nil {
		log.Fatal("cannot serve.", err)
	}
}
