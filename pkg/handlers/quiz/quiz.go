package quiz

import (
	"fmt"
	"github.com/dusnm/quiz/pkg/dto"
	"github.com/dusnm/quiz/pkg/services/password"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
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

func Generate(ctx *fiber.Ctx) error {
	quiz := new(dto.Quiz)

	if err := ctx.BodyParser(quiz); err != nil {
		return err
	}

	if quiz.Password != "" {
		p, err := password.New(quiz.Password)

		if err != nil {
			return err
		}

		//TODO: Write to DB
		fmt.Println(string(p.Hash))
	}

	//TODO: Write to DB
	uuid4 := utils.UUIDv4()

	return ctx.Redirect("/quiz/" + uuid4)
}
