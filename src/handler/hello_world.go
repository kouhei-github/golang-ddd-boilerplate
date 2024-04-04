package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

type IHelloWorldService interface {
	HelloMessage(id int) string
	WorldMessage(id int, userID int) string
}

type HelloWorldHandler struct {
	serviceHelloWorld IHelloWorldService
}

func NewHelloWorldHandler(
	serviceHelloWorld IHelloWorldService,
) *HelloWorldHandler {
	return &HelloWorldHandler{
		serviceHelloWorld: serviceHelloWorld,
	}
}

func (h *HelloWorldHandler) HelloWorld(c *fiber.Ctx) error {
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
