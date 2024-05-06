package handlers

import "github.com/gofiber/fiber/v2"

type LoginHandler interface {
	Login(c *fiber.Ctx) error
	SignUp(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}
