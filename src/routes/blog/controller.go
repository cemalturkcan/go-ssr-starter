package blog

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string
}

func getBySlug(c *fiber.Ctx) error {
	return c.Render("views/blog", fiber.Map{})
}

func Register(s *fiber.App) {
	group := s.Group("/blog")
	group.Get("/", getBySlug)
	group.Get("/:slug", getBySlug)
}
