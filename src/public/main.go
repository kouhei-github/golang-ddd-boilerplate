package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/golang-ddd-boboilerplate/config"
	"github.com/kouhei-github/golang-ddd-boboilerplate/di"
	"github.com/kouhei-github/golang-ddd-boboilerplate/providers"
)

func main() {
	env := config.NewConfigENV()
	env.EnvLoad()

	database := providers.NewDatabaseProvider()
	db, _, _ := database.Connect()
	//db.AutoMigrate(&models.User{}) // auto migrationできる
	if db == nil {
		panic("db is nil.")
	}

	app := fiber.New()

	group := app.Group("/api")
	r := di.NewRouter(*db)
	r.Register(group)
	if err := app.Listen(":8888"); err != nil {
		panic(err)
	}

}
