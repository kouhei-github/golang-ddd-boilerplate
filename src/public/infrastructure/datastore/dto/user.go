package dto

import (
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/models/user_models"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	UserName string `json:"user_name"`
	Email    string `gorm:"not null" gorm:"unique" json:"email"`
	Image    string `json:"image"`
	UserAuth UserAuth
}

type UserAuth struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	UserID       string `gorm:"not null" json:"user_id"`
	PasswordHash string `json:"-"`
	PasswordSalt string `json:"password_salt"`
}

func ToOrmUser(entity *user_models.User) (*User, error) {
	newUser := &User{
		UserName: string(entity.UserName),
		Email:    string(entity.Email),
		Image:    string(entity.Image),
		UserAuth: UserAuth{
			PasswordHash: string(entity.PasswordHash),
			PasswordSalt: string(entity.PasswordSalt),
		},
	}

	return newUser, nil
}

func ToEntityUser(newUser *User) (*user_models.User, error) {
	entityUser := user_models.User{
		ID: newUser.ID,
	}

	if newUser.Email != "" {
		email, err := user_models.NewEmail(newUser.Email)
		if err != nil {
			return nil, err
		}
		entityUser.Email = email
	}

	if newUser.UserName != "" {
		name, err := user_models.NewUserName(newUser.UserName)
		if err != nil {
			return nil, err
		}
		entityUser.UserName = name
	}

	if newUser.Image != "" {
		image, err := user_models.NewImage(newUser.Image)
		if err != nil {
			return nil, err
		}
		entityUser.Image = image
	}

	if newUser.UserAuth.PasswordSalt != "" {
		passwordSalt, err := user_models.NewPasswordSalt(newUser.UserAuth.PasswordSalt)
		if err != nil {
			return nil, err
		}
		entityUser.PasswordSalt = passwordSalt
	}

	if newUser.UserAuth.PasswordHash != "" {
		passwordHash, err := user_models.NewPasswordHash(newUser.UserAuth.PasswordHash)
		if err != nil {
			return nil, err
		}
		entityUser.PasswordHash = passwordHash
	}

	return &entityUser, nil
}
