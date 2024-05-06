package services

import (
	"github.com/kouhei-github/golang-ddd-boboilerplate/models"
	"github.com/kouhei-github/golang-ddd-boboilerplate/repositories"
)

type userService struct {
	user repositories.User
}

func NewUserService(user repositories.User) UserService {
	return &userService{
		user: user,
	}
}

func (u *userService) GetByID(id int) (*models.User, error) {
	user, err := u.user.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) GetByEmail(email string) (*models.User, error) {
	user, err := u.user.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) GetUserAuthByID(id int) (*models.UserAuth, error) {
	userAuth, err := u.user.GetUserAuthByID(id)
	if err != nil {
		return nil, err
	}
	return userAuth, nil
}

func (u *userService) UpdateUserAuth(userAuth *models.UserAuth) error {
	err := u.user.UpdateUserAuth(userAuth)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Create(user *models.User) error {
	err := u.user.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Update(user *models.User) error {
	err := u.user.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Delete(id int) error {
	err := u.user.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
