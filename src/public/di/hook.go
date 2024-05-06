// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/kouhei-github/golang-ddd-boboilerplate/handlers"
	"github.com/kouhei-github/golang-ddd-boboilerplate/repositories"
	"github.com/kouhei-github/golang-ddd-boboilerplate/routers"
	"github.com/kouhei-github/golang-ddd-boboilerplate/services"
	"gorm.io/gorm"
)

func NewRouter(db gorm.DB) routers.Router {
	// ユーザー
	user := repositories.NewUser(db)
	userService := services.NewUserService(user)
	authServic := services.NewAuthService()
	loginHandler := handlers.NewLoginHandler(userService, authServic)

	routerRouter := routers.NewRouter(loginHandler)
	return routerRouter
}
