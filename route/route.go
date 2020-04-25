package route

import (
	"../controller"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"net/http"
)

func Init() {
	app := fiber.New()

	app.Settings.Prefork = true

	config := cors.Config{AllowMethods: []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
	}}

	app.Use(cors.New(config))

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Post("/login", controller.Login)

	MemberRoute(app.Group("/members"))
	PlaceRoute(app.Group("/places"))
	ReviewRoute(app.Group("/reviews"))
	TourRoute(app.Group("/tours"))
	TranscriptRoute(app.Group("/transcripts"))

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
