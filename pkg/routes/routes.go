package routes

import (
	"github.com/dusnm/quiz/pkg/routes/home"
	"github.com/dusnm/quiz/pkg/routes/quiz"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	home.Routes(app)
	quiz.Routes(app)
}
