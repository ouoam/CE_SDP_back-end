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
				c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
				return
			}
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		result, err := model.GetPlace(id)
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
		place := new(model.Place)

		if err := c.BodyParser(&place); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := model.InsertPlace(place); err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := c.JSON(result); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
	})
}
