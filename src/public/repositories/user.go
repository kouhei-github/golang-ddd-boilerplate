package repositories

import (
	"errors"
	"github.com/kouhei-github/golang-ddd-boboilerplate/models"

	"gorm.io/gorm"
)

type user struct {
	db gorm.DB
}

func NewUser(db gorm.DB) User {
	return &user{
		db: db,
	}
}

func (u *user) GetByID(id int) (*models.User, error) {
	user := models.User{}
	err := u.db.Find(&user, id).Error
	if err != nil {
		return nil, err
	}
	if user.ID != id {
		return nil, errors.New("not found")
	}
	return &user, nil
}

func (u *user) GetByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := u.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Email != email {
		return nil, errors.New("not found")
	}
	return &user, nil
}

func (u *user) GetUserAuthByID(id int) (*models.UserAuth, error) {
	userAuth := models.UserAuth{}
	err := u.db.Where("user_id = ?", id).Find(&userAuth).Error
	if err != nil {
		return nil, err
	}

	return &userAuth, nil
}

func (u *user) UpdateUserAuth(userAuth *models.UserAuth) error {
	err := u.db.Save(&userAuth).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *user) Create(user *models.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *user) Update(user *models.User) error {
	if user.ID == 0 {
		return errors.New("id is not valid")
	}

	return u.db.Updates(&user).Error
}

func (u *user) Delete(id int) error {
	if id == 0 {
		return errors.New("id is not valid")
	}

	return u.db.Delete(&models.User{}, id).Error
}
