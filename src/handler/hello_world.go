package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/golang-ddd-boboilerplate/service"
	"strconv"
	"strings"
)

type HelloWorldHandler interface {
	HelloWorld(c *fiber.Ctx) error
}

type helloWorld struct {
	serviceHelloWorld service.HelloWorldService
}

func NewHelloWorldHandler(
	serviceHelloWorld service.HelloWorldService,
) HelloWorldHandler {
	return &helloWorld{
		serviceHelloWorld: serviceHelloWorld,
	}
}

func (h *helloWorld) HelloWorld(c *fiber.Ctx) error {
	query := c.Query("id")
	userId, err := strconv.Atoi(query)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).SendString("Internal Server Error")
	}
	hello := h.serviceHelloWorld.HelloMessage(userId)

	word := h.serviceHelloWorld.WorldMessage(userId, userId)
	if err != nil {
		fmt.Println(err.Error())

		return c.Status(500).SendString("Internal Server Error")
	}
	type response struct {
		Message string `json:"message" example:"1"`
	}

	helloWorldMessage := strings.ToUpper(hello) + " " + strings.ToUpper(word)
	res := response{
		Message: helloWorldMessage,
	}
	return c.Status(200).JSON(res)
}
