package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
	"net/http"
)

func TourRoute(route *fiber.Group) {
	route.Get("/:id", parseIntParams("id"), func(c *fiber.Ctx) {
		tour := new(model.TourDetailSearch)
		tour.ID.SetValid(c.Locals("params_id").(int64))
		controller.Get(c, tour)
	})

	route.Get("/:id/reviews", parseIntParams("id"), func(c *fiber.Ctx) {
		review := new(model.ReviewWithTour)
		controller.List(c, review, c.Locals("params_id").(int64))
	})

	route.Get("/:id/lists", parseIntParams("id"), func(c *fiber.Ctx) {
		list := new(model.ListWithTour)
		controller.List(c, list, c.Locals("params_id").(int64))
	})

	route.Post("/:id/lists", parseIntParams("id"), func(c *fiber.Ctx) {
		listBody := new(model.ListUpdateBody)
		if err := c.BodyParser(listBody); err != nil {
			_ = c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}
		list := new(model.ListUpdate)
		controller.Get(c, list, c.Locals("params_id").(int64), listBody.Place)
	})

	route.Get("/:id/transcripts", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		transcript := new(model.TranscriptWithTour)
		controller.List(c, transcript, c.Locals("params_id").(int64))
	})

	route.Get("/:id/transcripts/:user", controller.CheckLogin, parseIntParams("id", "user"), func(c *fiber.Ctx) {
		transcript := new(model.TranscriptWithTour)
		transcript.User.SetValid(c.Locals("params_user").(int64))
		controller.Get(c, transcript, c.Locals("params_id").(int64))
	})

	route.Post("/:id/transcripts/:user", controller.CheckLogin, parseIntParams("id", "user"), func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		transcript.Tour.SetValid(c.Locals("params_id").(int64))
		transcript.User.SetValid(c.Locals("params_user").(int64))
		controller.Update(c, transcript)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		tour := new(model.Tour)
		tour.Owner.SetValid(c.Locals("user_id").(int64))
		controller.New(c, tour)
	})

	route.Put("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		// todo update own
		tour := new(model.Tour)
		tour.ID.SetValid(c.Locals("params_id").(int64))
		tour.Owner.SetValid(c.Locals("user_id").(int64))
		controller.Update(c, tour)
	})

	route.Get("/", func(c *fiber.Ctx) {
		tour := new(model.TourDetailSearch)
		controller.List(c, tour)
	})
}
