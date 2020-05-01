package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func FavoriteRoute(route *fiber.Group) {
	route.Get("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		favorite := new(model.FavoriteWithUser)
		favorite.Tour.SetValid(c.Locals("params_id").(int64))
		controller.Get(c, favorite, c.Locals("user_id").(int64))
	})

	route.Delete("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		favorite := new(model.Favorite)
		favorite.Tour.SetValid(c.Locals("params_id").(int64))
		favorite.User.SetValid(c.Locals("user_id").(int64))
		controller.Delete(c, favorite)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		favorite := new(model.Favorite)
		favorite.User.SetValid(c.Locals("user_id").(int64))
		controller.New(c, favorite)
	})

	route.Get("/", controller.CheckLogin, func(c *fiber.Ctx) {
		favorite := new(model.FavoriteWithUser)
		controller.List(c, favorite, c.Locals("user_id").(int64))
	})
}
