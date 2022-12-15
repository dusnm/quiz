package quiz

import (
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		return ctx.Redirect("/")
	}

	//TODO: Query the DB for a quiz with this id
	// check if it's locked and then render either make or play
	return ctx.Render("views/quiz/make", fiber.Map{})
}
