package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/kouhei-github/golang-ddd-boboilerplate/config"
)

type authService struct {
	secretKey []byte
}

func NewAuthService() AuthService {
	return &authService{
		secretKey: config.JWT_SECRET_KEY,
	}
}

type jwtCustomClaims struct {
	userClaim
	jwt.StandardClaims
}

type userClaim struct {
	UserID    int
	UserName  string
	UserEmail string
}

func (s *authService) GetClaimFromToken(token string) (*jwtCustomClaims, error) {
	claim := jwtCustomClaims{}
	result, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !result.Valid {
		return nil, errors.New("token is not valid")
	}

	return &claim, nil
}
