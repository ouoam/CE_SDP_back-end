package route

import (
	"../model"
	"github.com/gofiber/fiber"
	"net/http"
	"strconv"
	"strings"
)

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

func GetID(c *fiber.Ctx, dataModel model.WithID) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		if numError, ok := err.(*strconv.NumError); ok {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
			return
		}
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	dataModel.SetID(id)

	if err := dataModel.GetDB(); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			_ = c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Place not found"})
			return
		}
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func Post(c *fiber.Ctx, dataModel model.WithID) {
	if err := c.BodyParser(&dataModel); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if err := dataModel.AddDB(); err != nil {
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

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}

func PutID(c *fiber.Ctx, dataModel model.WithID) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		if numError, ok := err.(*strconv.NumError); ok {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": numError.Err.Error()})
			return
		}
		_ = c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if err := c.BodyParser(&dataModel); err != nil {
		_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	dataModel.SetID(id)

	if err := dataModel.UpdateDB(); err != nil {
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

	if err := c.JSON(dataModel); err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
}