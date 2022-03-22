package main

import (
	"github.com/gofiber/fiber/v2"
)

func initQuizRoutes(a *fiber.App) {

	a.Get("/quiz/create", createQuiz)

}

func createQuiz(c *fiber.Ctx) error {
	if err := c.SendString("ping!"); err != nil {
		return err
	}
	return nil
}
