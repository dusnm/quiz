package home

import "github.com/gofiber/fiber/v2"

func Index(ctx *fiber.Ctx) error {
	return ctx.Render("views/home/home", fiber.Map{})
}
