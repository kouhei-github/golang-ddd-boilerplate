package repository

import (
	"gorm.io/gorm"
)

type HelloWorld interface {
	GetHello(id int) (string, error)
	GetWorld() string
}

type helloWorld struct {
	db gorm.DB
}

func NewHelloWorldRepository(db gorm.DB) HelloWorld {
	return &helloWorld{
		db: db,
	}
}

func (c *helloWorld) GetHello(id int) (string, error) {
	hello := "World"
	return hello, nil
}

func (c *helloWorld) GetWorld() string {
	world := "World"
	return world
}
