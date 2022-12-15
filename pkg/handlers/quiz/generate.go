package quiz

import (
	"fmt"
	"github.com/dusnm/quiz/pkg/dto"
	"github.com/dusnm/quiz/pkg/services/password"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

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
