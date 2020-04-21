package route

import (
	"github.com/gofiber/fiber"
	"net/http"
	"strconv"
	"strings"

	"../model"
)

func PlaceRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			if numError, ok := err.(*strconv.NumError); ok {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
				return
			}
			_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		result, err := model.GetPlace(id)
		if err != nil {
			if strings.Contains(err.Error(), "no rows") {
				_ = c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Place not found"})
				return
			}
			_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.JSON(result); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})

	route.Post("/", func(c *fiber.Ctx) {
		place := new(model.Place)

		if err := c.BodyParser(&place); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := model.AddPlace(place); err != nil {
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

		if err := c.JSON(place); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		place := new(model.Place)

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			if numError, ok := err.(*strconv.NumError); ok {
				_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
				return
			}
			_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.BodyParser(&place); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := model.UpdatePlace(place, id); err != nil {
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

		if err := c.JSON(place); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})
}
