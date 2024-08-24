package auth_use_case

import (
	"errors"
	"github.com/kouhei-github/golang-ddd-boboilerplate/application/use_case/impluments/auth_use_case_imp"
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/interface/external"
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/interface/repositories"
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/models/user_models"
)

type LoginUseCase struct {
	ur  repositories.UserRepository
	jte external.JWTTokenExternal
}

func NewLoginUseCase(
	userRepo repositories.UserRepository,
	jte external.JWTTokenExternal,
) LoginUseCase {
	return LoginUseCase{ur: userRepo, jte: jte}
}

func (lu LoginUseCase) Execute(email, password string) (*auth_use_case_imp.LoginResponse, error) {
	emailVo, err := user_models.NewEmail(email)
	if err != nil {
		return nil, err
	}
	passwordVo, err := user_models.NewPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := lu.ur.GetByEmail(string(emailVo))
	if err != nil {
		return nil, err
	}

	authUser, err := lu.ur.GetUserAuthByID(user.ID)
	if err != nil {
		return nil, err
	}
	if !authUser.CheckPassword(passwordVo.Native) {
		return nil, errors.New("パスワードが一致しません。")
	}

	// Create token
	accessToken, err := lu.jte.GenerateToken(AccessTokenExpires, user.ID, string(user.UserName), string(user.Email))
	if err != nil {
		return nil, err
	}

	refreshToken, err := lu.jte.GenerateToken(RefreshTokenExpires, user.ID, string(user.UserName), string(user.Email))
	if err != nil {
		return nil, err
	}

	return &auth_use_case_imp.LoginResponse{
		UserId:             user.ID,
		Token:              accessToken,
		AccessTokenExpires: int(AccessTokenExpires.Seconds()),
		RefreshToken:       refreshToken,
		AvatarURL:          string(user.Image),
	}, nil

}
