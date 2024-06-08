package repositories

import (
	"errors"
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/models/user_models"
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/repositories"
	"github.com/kouhei-github/golang-ddd-boboilerplate/infrastructure/datastore/dto"

	"gorm.io/gorm"
)

type user struct {
	db gorm.DB
}

func NewUserRepository(db gorm.DB) repositories.UserRepository {
	return &user{
		db: db,
	}
}

func (u *user) GetByID(id int) (*user_models.User, error) {
	user := dto.User{}
	err := u.db.Find(&user, id).Error
	if err != nil {
		return nil, err
	}
	if user.ID != id {
		return nil, errors.New("not found")
	}
	entityUser, err := dto.ToEntityUser(&user)
	if err != nil {
		return nil, errors.New("not found")
	}
	return entityUser, nil
}

func (u *user) GetByEmail(email string) (*user_models.User, error) {
	user := dto.User{}
	err := u.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	entityUser, err := dto.ToEntityUser(&user)
	if err != nil {
		return nil, errors.New("not found")
	}
	return entityUser, nil
}

func (u *user) GetUserAuthByID(id int) (*user_models.User, error) {
	var user dto.User
	err := u.db.Preload("UserAuth").Find(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	entityUser, err := dto.ToEntityUser(&user)
	if err != nil {
		return nil, errors.New("not found")
	}
	return entityUser, nil
}

//func (u *user) UpdateUserAuth(userAuth *models.UserAuth) error {
//	err := u.db.Save(&userAuth).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (u *user) Create(entUser *user_models.User) error {
	// entityからORMに変換
	userOrm, err := dto.ToOrmUser(entUser)
	if err != nil {
		return err
	}

	// 保存する
	if err = u.db.Create(userOrm).Error; err != nil {
		return err
	}

	return nil
}

//func (u *user) Update(user *models.User) error {
//	if user.ID == 0 {
//		return errors.New("id is not valid")
//	}
//
//	return u.db.Updates(&user).Error
//}
//
//func (u *user) Delete(id int) error {
//	if id == 0 {
//		return errors.New("id is not valid")
//	}
//
//	return u.db.Delete(&models.User{}, id).Error
//}
