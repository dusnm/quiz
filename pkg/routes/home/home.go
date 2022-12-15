package home

import (
	"github.com/dusnm/quiz/pkg/handlers/home"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", home.Index)
}
