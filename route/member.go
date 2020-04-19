package route

import "github.com/gofiber/fiber"

func MemberRoute(route *fiber.Group) {
	route.Get("/:id", func(c *fiber.Ctx) {
		c.Send("Hello, World!" + c.Params("id"))
	})
}
