package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/golang-ddd-boboilerplate/config"
	"github.com/kouhei-github/golang-ddd-boboilerplate/di"
	"github.com/kouhei-github/golang-ddd-boboilerplate/provider"
)

func main() {
	env := config.NewConfigENV()
	env.EnvLoad()

	database := provider.NewDatabaseProvider()
	db, _, _ := database.Connect()
	if db == nil {
		panic("db is nil.")
	}

	app := fiber.New()

	group := app.Group("")
	r := di.NewRouter(*db)
	r.Register(group)

	if err := app.Listen(":8888"); err != nil {
		panic(err)
	}

}
