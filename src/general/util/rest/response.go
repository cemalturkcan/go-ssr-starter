package rest

import "github.com/gofiber/fiber/v2"

func jsonResponse(c *fiber.Ctx, code string, message string) error {
	return c.JSON(&fiber.Map{
		"code":    code,
		"message": message,
	})
}

func ErrorRes(c *fiber.Ctx, code string, message string) error {
	return jsonResponse(c, code, message)
}
