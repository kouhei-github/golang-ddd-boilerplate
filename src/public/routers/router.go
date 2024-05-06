package routers

import (
	"github.com/gofiber/fiber/v2"
)

type IHelloWorldHandler interface {
	HelloWorld(c *fiber.Ctx) error
}

type Router struct {
	helloWorld IHelloWorldHandler
}

func NewRouter(
	helloWorld IHelloWorldHandler,
) *Router {
	return &Router{
		helloWorld: helloWorld,
	}
}

func (r *Router) Register(c fiber.Router) {
	common := c.Group("/api")
	common.Get("/test", r.helloWorld.HelloWorld)
}
