package auth_use_case

import (
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/interface/external"
	"github.com/kouhei-github/golang-ddd-boboilerplate/domain/interface/repositories"
	"os"
	"time"
)

type RefreshTokenUseCase struct {
	ur  repositories.UserRepository
	jte external.JWTTokenExternal
}

func NewRefreshTokenUseCase(
	userRepo repositories.UserRepository,
	jwtExternal external.JWTTokenExternal,
) RefreshTokenUseCase {
	return RefreshTokenUseCase{ur: userRepo, jte: jwtExternal}
}

var unit = time.Minute

var (
	JWT_SECRET_KEY      = []byte(os.Getenv("JWT_SECRET_KEY"))
	RefreshTokenExpires = 60 * 24 * 30 * 2 * unit
	AccessTokenExpires  = 20 * unit
)

type Response struct {
	UserId             int    `json:"user_id"`
	Token              string `json:"access_token"`
	AccessTokenExpires int    `json:"access_token_expires"`
	RefreshToken       string `json:"refreshToken"`
	ImageURL           string `json:"image_url"`
}

func (ru RefreshTokenUseCase) Execute(refreshToken string) (*Response, error) {
	userId, name, email, err := ru.jte.GetClaimFromToken(refreshToken, string(JWT_SECRET_KEY))

	if err != nil {
		return nil, err
	}

	user, err := ru.ur.GetByID(userId)
	if err != nil {
		return nil, err
	}

	// Create token
	accessToken, err := ru.jte.GenerateToken(AccessTokenExpires, userId, name, email)
	if err != nil {
		return nil, err
	}

	refreshToken, err = ru.jte.GenerateToken(RefreshTokenExpires, userId, name, email)
	if err != nil {
		return nil, err
	}

	res := Response{
		UserId:             user.ID,
		Token:              accessToken,
		AccessTokenExpires: int(AccessTokenExpires.Seconds()),
		RefreshToken:       refreshToken,
		ImageURL:           string(user.Image),
	}

	return &res, err

}
