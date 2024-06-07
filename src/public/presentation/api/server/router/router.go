package router

import (
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Register(c fiber.Router)
}

type router struct {
	auth AuthHandler
}

func NewRouter(auth AuthHandler) Router {
	return &router{
		auth: auth,
	}
}

func (r *router) Register(c fiber.Router) {
	common := c.Group("/v1")
	common.Post("/login", r.auth.Login)
	common.Post("/refresh", r.auth.RefreshToken)
	common.Post("/signup", r.auth.SignUp)
}
