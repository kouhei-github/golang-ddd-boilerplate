package repositories

import "github.com/kouhei-github/golang-ddd-boboilerplate/models"

type User interface {
	GetByID(id int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetUserAuthByID(id int) (*models.UserAuth, error)
	UpdateUserAuth(userAuth *models.UserAuth) error
	Create(u *models.User) error
	Update(u *models.User) error
	Delete(id int) error
}
