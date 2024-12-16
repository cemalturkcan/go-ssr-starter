package index

import (
	"github.com/gofiber/fiber/v2"
)

func get(c *fiber.Ctx) error {
	return c.Render("views/filter", fiber.Map{})
}

func Register(s *fiber.App) {
	group := s.Group("/")
	group.Get("/", get)
}
