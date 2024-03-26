package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/golang-ddd-boboilerplate/handler"
)

type Router interface {
	Register(c fiber.Router)
}

type router struct {
	helloWorld handler.HelloWorldHandler
}

func NewRouter(
	helloWorld handler.HelloWorldHandler,
) Router {
	return &router{
		helloWorld: helloWorld,
	}
}

func (r router) Register(c fiber.Router) {
	common := c.Group("/api")
	common.Get("/test", r.helloWorld.HelloWorld)
}
