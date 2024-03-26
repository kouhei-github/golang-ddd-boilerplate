package service

import (
	"github.com/kouhei-github/golang-ddd-boboilerplate/repository"
)

type HelloWorldService interface {
	HelloMessage(id int) string
	WorldMessage(id int, userID int) string
}

type helloWorldService struct {
	repositoryHelloWorld repository.HelloWorld
}

func NewHelloWorldService(
	repositoryHelloWorld repository.HelloWorld,
) HelloWorldService {
	return &helloWorldService{
		repositoryHelloWorld: repositoryHelloWorld,
	}
}

func (c *helloWorldService) HelloMessage(tid int) string {
	msg, _ := c.repositoryHelloWorld.GetHello(tid)
	return msg
}

func (c *helloWorldService) WorldMessage(id int, userID int) string {
	return c.repositoryHelloWorld.GetWorld()
}
