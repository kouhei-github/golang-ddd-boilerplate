package auth_use_case

import (
	"errors"
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/models/user_models"
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/repositories"
)

type SignUpUseCase struct {
	ur repositories.UserRepository
}

func NewSignUpUseCase(
	userRepo repositories.UserRepository,
) SignUpUseCase {
	return SignUpUseCase{ur: userRepo}
}

func (su SignUpUseCase) Execute(email, password string) error {
	emailVo, err := user_models.NewEmail(email)
	if err != nil {
		return err
	}
	passwordVo, err := user_models.NewPassword(&password)
	if err != nil {
		return err
	}

	user, err := su.ur.GetByEmail(string(emailVo))
	if err != nil {
		return err
	}
	if user.Email != "" {
		return errors.New("ユーザは既に存在します。")
	}

	newUser, err := user_models.NewUser(string(emailVo), string(passwordVo), "", "", "", "")
	if err != nil {
		return err
	}

	newUser.SetPassword(newUser.Password)

	if err := su.ur.Create(newUser); err != nil {
		return err
	}

	return nil
}
