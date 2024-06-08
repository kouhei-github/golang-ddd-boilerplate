package repositories

import (
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/models/user_models"
)

type UserRepository interface {
	GetByID(id int) (*user_models.User, error)
	GetByEmail(email string) (*user_models.User, error)
	GetUserAuthByID(id int) (*user_models.User, error)
	//UpdateUserAuth(userAuth *models.UserAuth) error
	Create(email user_models.Email, password user_models.Password) error
	//Update(u *models.User) error
	//Delete(id int) error
}
