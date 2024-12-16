package filter

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string
}

func getByCity(c *fiber.Ctx) error {
	return c.Render("views/filter", fiber.Map{
		"VipUsers": []User{
			{Name: "Alice"},
			{Name: "Ali"},
		},
	})
}

func getByCityAndDistrict(c *fiber.Ctx) error {
	return c.Render("views/filter", fiber.Map{
		"VipUsers": []User{
			{Name: "Alice"},
			{Name: "Ali"},
		},
	}, "layouts/main")
}

func Register(s *fiber.App) {
	group := s.Group("/filter")
	group.Get("/:city", getByCity)
	group.Get("/:city/:district", getByCityAndDistrict)
}
