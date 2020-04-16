package route

import "github.com/gofiber/fiber"

func Init() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Listen(3000)
}
