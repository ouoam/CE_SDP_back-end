package route

import (
	"../controller"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"net/http"
	"strconv"
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
	app.Post("/forgot", controller.ForgotPassword)

	MemberRoute(app.Group("/members"))
	PlaceRoute(app.Group("/places"))
	ReviewRoute(app.Group("/reviews"))
	TourRoute(app.Group("/tours"))
	TranscriptRoute(app.Group("/transcripts"))
	MessageRoute(app.Group("/messages"))

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}

func parseIntParams(params... string) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		for _, param := range params {
			id, err := strconv.ParseInt(c.Params(param), 10, 64)
			if err != nil {
				if numError, ok := err.(*strconv.NumError); ok {
					_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
					return
				}
				_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
				return
			}
			c.Locals("params_" + param, id)
		}
		c.Next()
		return
	}
}