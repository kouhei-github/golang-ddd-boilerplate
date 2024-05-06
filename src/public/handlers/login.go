package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/golang-ddd-boboilerplate/config"
	"github.com/kouhei-github/golang-ddd-boboilerplate/models"
	"github.com/kouhei-github/golang-ddd-boboilerplate/services"
	"net/http"
)

type loginHandler struct {
	user services.UserService
	auth services.AuthService
}

func NewLoginHandler(
	user services.UserService,
	auth services.AuthService,
) LoginHandler {
	return &loginHandler{
		user: user,
		auth: auth,
	}
}

func (h *loginHandler) Login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.user.GetByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	userAuth, err := h.user.GetUserAuthByID(int(user.ID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if !userAuth.CheckPassword(req.Password) {
		return c.Status(fiber.StatusBadRequest).JSON("password is not correct")
	}

	// Create token
	accessToken, err := user.GenerateToken(config.AccessTokenExpires)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	refreshToken, err := user.GenerateToken(config.RefreshTokenExpires)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	type response struct {
		UserId             int    `json:"userID"`
		Token              string `json:"accessToken"`
		AccessTokenExpires int    `json:"accessTokenExpires"`
		RefreshToken       string `json:"refreshToken"`
		AvatarURL          string `json:"avatarURL"`
		Role               int    `json:"role"`
	}

	res := response{
		UserId:             user.ID,
		Token:              accessToken,
		AccessTokenExpires: int(config.AccessTokenExpires.Seconds()),
		RefreshToken:       refreshToken,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *loginHandler) SignUp(c *fiber.Ctx) error {
	type request struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("hear")
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	newUser := models.User{
		UserName: req.Email,
		Email:    req.Email,
	}

	newUserAuth := models.UserAuth{}
	newUserAuth.SetPassword(req.Password)

	if temUser, err := h.user.GetByEmail(newUser.Email); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fmt.Sprintf("user already exists: %v", temUser.Email))
	}

	newUser.UserAuth = newUserAuth
	if err := h.user.Create(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.JSON(http.StatusOK, "success")
}

func (h *loginHandler) RefreshToken(c *fiber.Ctx) error {
	type request struct {
		RefreshToken string `json:"refreshToken"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	claim, err := h.auth.GetClaimFromToken(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.user.GetByID(claim.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Create token
	accessToken, err := user.GenerateToken(config.AccessTokenExpires)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	refreshToken, err := user.GenerateToken(config.RefreshTokenExpires)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	type response struct {
		UserId             int    `json:"userID"`
		Token              string `json:"accessToken"`
		AccessTokenExpires int    `json:"accessTokenExpires"`
		RefreshToken       string `json:"refreshToken"`
		ImageURL           string `json:"avatarURL"`
	}

	imageURL := ""
	if user.Image != "" {
		imageURL = user.Image
	}

	res := response{
		UserId:             user.ID,
		Token:              accessToken,
		AccessTokenExpires: int(config.AccessTokenExpires.Seconds()),
		RefreshToken:       refreshToken,
		ImageURL:           imageURL,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
