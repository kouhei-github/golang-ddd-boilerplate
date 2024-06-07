package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/golang-ddd-boboilerplate/di"
	"github.com/kouhei-github/golang-ddd-boboilerplate/env"
	"github.com/kouhei-github/golang-ddd-boboilerplate/infrastructure/datastore"
)

func main() {
	envs := env.NewConfigENV()
	envs.EnvLoad()

	// 環境変数を型に変換
	el := env.NewLib()
	if err := el.CheckValues(); err != nil {
		panic("Environment variables not set: " + err.Error())
	}

	// databaseの初期化
	database := datastore.NewDatabaseProvider(el)
	db, _, _ := database.Connect()
	//db.AutoMigrate(&dto.User{}) // auto migrationできる
	if db == nil {
		panic("db is nil.")
	}

	// fiberの起動
	app := fiber.New()

	group := app.Group("/api")
	r := di.NewRouter(*db, el)
	r.Register(group)
	if err := app.Listen(":8888"); err != nil {
		panic(err)
	}

}
