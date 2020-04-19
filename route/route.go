package route

import "github.com/gofiber/fiber"

func Init() {
	app := fiber.New()

	app.Settings.Prefork = true

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	MemberRoute(app.Group("/members"))
	PlaceRoute(app.Group("/places"))

	app.Listen(3000)
}
