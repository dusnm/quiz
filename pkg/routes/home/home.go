package home

import "github.com/gofiber/fiber/v2"

func home(ctx *fiber.Ctx) error {
	return ctx.Render("views/home", fiber.Map{})
}

func Routes(app *fiber.App) {
	app.Get("/", home)
}
