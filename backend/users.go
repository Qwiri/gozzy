package main

import (
	"github.com/gofiber/fiber/v2"
)

func initUserRoutes(a *fiber.App) {

	a.Get("/users/login", login)

}

func login(c *fiber.Ctx) error {
	if err := c.SendString("ping!"); err != nil {
		return err
	}
	return nil
}
