package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/golang-ddd-boboilerplate/handlers"
)

type Router interface {
	Register(c fiber.Router)
}

type router struct {
	login handlers.LoginHandler
}

func NewRouter(login handlers.LoginHandler) Router {
	return &router{
		login: login,
	}
}

func (r *router) Register(c fiber.Router) {
	common := c.Group("/v1")
	common.Post("/login", r.login.Login)
	common.Post("/refresh", r.login.RefreshToken)
	common.Post("/signup", r.login.SignUp)
}
