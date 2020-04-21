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
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			if numError, ok := err.(*strconv.NumError); ok {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
				return
			}
			_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		member := new(model.Member)
		member.ID.SetValid(id)

		if err := member.GetDB(); err != nil {
			if strings.Contains(err.Error(), "no rows") {
				_ = c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Place not found"})
				return
			}
			_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.JSON(member); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})

	route.Post("/", func(c *fiber.Ctx) {
		member := new(model.Member)

		if err := c.BodyParser(&member); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := member.AddDB(); err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				if strings.Contains(err.Error(), "username") {
					_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this username."})
					return
				}
				if strings.Contains(err.Error(), "email") {
					_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this E-Mail."})
					return
				}
			}
			_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.JSON(member); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		member := new(model.Member)

		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			if numError, ok := err.(*strconv.NumError); ok {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
				return
			}
			_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.BodyParser(&member); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}

		member.ID.SetValid(id)

		if err := member.UpdateDB(); err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				if strings.Contains(err.Error(), "username") {
					_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this username."})
					return
				}
				if strings.Contains(err.Error(), "email") {
					_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Other member already used this E-Mail."})
					return
				}
			}
			_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.JSON(member); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})
}
