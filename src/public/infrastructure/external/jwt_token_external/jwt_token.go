package jwt_token_external

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/kouhei-github/golang-ddd-boboilerplate/application/use_case/external"
	"os"
	"time"
)

var (
	JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
)

type jwtCustomClaims struct {
	userClaim
	jwt.StandardClaims
}

type userClaim struct {
	userID    int
	userName  string
	userEmail string
}

func NewJwtToken() external.JWTTokenExternal {
	return jwtCustomClaims{}
}

func (jcc jwtCustomClaims) GenerateToken(expiredTime time.Duration, userId int, name, email string) (string, error) {
	// Default: Standard Claim, Custom: User Clamを用いる
	SECRET_KEY := []byte(JWT_SECRET_KEY)
	claims := &jwtCustomClaims{
		userClaim{
			userID:    userId,
			userName:  name,
			userEmail: email,
		},
		jwt.StandardClaims{
			Audience:  fmt.Sprintf("%v.test.com", userId),                       // 観測者のUID
			ExpiresAt: time.Now().Add(expiredTime).Unix(),                       // 有効期限14日
			Id:        fmt.Sprintf("%v.test.com/%v", userId, time.Now().Unix()), // 発行ごとのID
			IssuedAt:  time.Now().Unix(),                                        // 発行日時
			Issuer:    "go.test.com",                                            // システムの名前
			NotBefore: time.Now().Unix(),                                        // いつから可能か
			Subject:   fmt.Sprintf("%v.test.com", userId),                       // 全世界でユニーク
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token.Claims.Valid() != nil {
		return "", token.Claims.Valid()
	}
	jwt, err := token.SignedString(SECRET_KEY)
	return jwt, err
}

func (jcc jwtCustomClaims) GetClaimFromToken(token string, secretKey string) (userId int, name, email string, err error) {
	userId, email, name = 0, "", ""
	claim := jwtCustomClaims{}
	result, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return
	}
	err = errors.New("token is not valid")
	if !result.Valid {
		return
	}

	userId = claim.userID
	name = claim.userName
	email = claim.userEmail
	return
}
