package route

import (
	"github.com/gofiber/fiber"
	"net/http"
	"strconv"
	"strings"

	"../model"
)

func MemberRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			if numError, ok := err.(*strconv.NumError); ok {
				c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
				return
			}
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		result, err := model.GetMember(id)
		if err != nil {
			if strings.Contains(err.Error(), "no rows") {
				c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Place not found"})
				return
			}
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.JSON(result); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})

	route.Post("/", func(c *fiber.Ctx) {
		member := new(model.Member)

		if err := c.BodyParser(&member); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := model.AddMember(member); err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				if strings.Contains(err.Error(), "username") {
					c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this username."})
					return
				}
				if strings.Contains(err.Error(), "email") {
					c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this E-Mail."})
					return
				}
			}
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.JSON(member); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})
}
