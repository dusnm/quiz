package quiz

import (
	"github.com/dusnm/quiz/pkg/handlers/quiz"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	quizRoutes := app.Group("/quiz")

	quizRoutes.Get("/:id", quiz.Index)
	quizRoutes.Post("/generate", quiz.Generate)
}
