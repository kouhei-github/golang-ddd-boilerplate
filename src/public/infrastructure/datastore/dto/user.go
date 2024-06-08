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

func EntityUserToOrmUser(entity *user_models.User) (*User, error) {
	newUser := &User{
		UserName: entity.UserName,
		Email:    entity.Email,
		Image:    entity.Image,
		UserAuth: UserAuth{
			PasswordHash: entity.PasswordHash,
			PasswordSalt: entity.PasswordSalt,
		},
	}

	return newUser, nil
}

func OrmUserToEntityUser(newUser *User) (*user_models.User, error) {

	original := &user_models.User{
		ID:           newUser.ID,
		UserName:     newUser.UserName,
		Email:        newUser.Email,
		Image:        newUser.Image,
		PasswordHash: newUser.UserAuth.PasswordHash,
		PasswordSalt: newUser.UserAuth.PasswordSalt,
	}

	return original, nil
}
