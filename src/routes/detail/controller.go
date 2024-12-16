package detail

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string
}

func get(c *fiber.Ctx) error {
	return c.Render("views/user-detail", fiber.Map{
		"Slug": c.Params("slug"),
	})
}

func Register(s *fiber.App) {
	group := s.Group("/detail")
	group.Get("/:slug", get)
}
