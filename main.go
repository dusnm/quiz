package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"log"
	"net/http"
)

//go:embed views/*
var views embed.FS

//go:embed static/*
var staticAssets embed.FS

func main() {
	engine := html.NewFileSystem(http.FS(views), ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/layouts/main",
	})

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(staticAssets),
		PathPrefix: "static",
		Browse:     false,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/home", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}
