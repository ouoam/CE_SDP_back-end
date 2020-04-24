package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
	"net/http"
)

func MemberRoute(route *fiber.Group) {
	route.Get("/:id", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.GetID(c, member)
	})

	route.Post("/", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.Post(c, member)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.PutID(c, member)
	})

	route.Get("/", func(c *fiber.Ctx) {
		member := new(model.Member)
		if err := c.BodyParser(&member); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}

		members, err := member.ListDB()
		if err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		_ = c.JSON(members)
	})
}
