package repositories

import (
	"gorm.io/gorm"
)

type HelloWorld struct {
	db gorm.DB
}

func NewHelloWorldRepository(db gorm.DB) *HelloWorld {
	return &HelloWorld{
		db: db,
	}
}

func (c *HelloWorld) GetHello(id int) (string, error) {
	hello := "World"
	return hello, nil
}

func (c *HelloWorld) GetWorld() string {
	world := "World"
	return world
}
