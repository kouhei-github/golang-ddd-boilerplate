package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/golang-ddd-boboilerplate/application/use_case/auth_use_case"
	"github.com/kouhei-github/golang-ddd-boboilerplate/presentation/api/server/router"
)

type authHandler struct {
	su auth_use_case.SignUpUseCase
	ru auth_use_case.RefreshTokenUseCase
	lu auth_use_case.LoginUseCase
}

func NewAuthHandler(
	su auth_use_case.SignUpUseCase,
	ru auth_use_case.RefreshTokenUseCase,
	lu auth_use_case.LoginUseCase,
) router.AuthHandler {
	return authHandler{su, ru, lu}
}

func (h authHandler) SignUp(c *fiber.Ctx) error {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req request
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("hear")
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// UseCaseの呼び出し
	if err := h.su.Execute(req.Email, req.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("done")
}

func (h authHandler) RefreshToken(c *fiber.Ctx) error {
	type request struct {
		RefreshToken string `json:"refreshToken"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	response, err := h.ru.Execute(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h authHandler) Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	response, err := h.lu.Execute(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
