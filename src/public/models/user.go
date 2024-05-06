package models

import (
	"fmt"
	config2 "github.com/kouhei-github/golang-ddd-boboilerplate/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	UserName string `json:"user_name"`
	Email    string `gorm:"not null" gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Image    string `json:"image"`
	UserAuth UserAuth
}

type UserAuth struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	PasswordHash string `json:"-"`
	PasswordSalt string `json:"password_salt"`
	UserID       int    `json:"user_id"`
}

func (u *UserAuth) SetPassword(password string) {
	saltLength := 16
	salt, _ := config2.GenerateSalt(saltLength)
	u.PasswordSalt = salt
	hash, _ := config2.HashPassword(password, []byte(salt))
	u.PasswordHash = hash
}

func (u *UserAuth) CheckPassword(password string) bool {
	return config2.ComparePasswords(u.PasswordHash, password, []byte(u.PasswordSalt))
}

type jwtCustomClaims struct {
	UserClaim
	jwt.StandardClaims
}

type UserClaim struct {
	UserID    int
	UserName  string
	UserEmail string
}

func (u *User) GenerateToken(expiredTime time.Duration) (string, error) {
	// Default: Standard Claim, Custom: User Clamを用いる
	SECRET_KEY := []byte(config2.JWT_SECRET_KEY)
	claims := &jwtCustomClaims{
		UserClaim{
			UserID:    u.ID,
			UserName:  u.UserName,
			UserEmail: u.Email,
		},
		jwt.StandardClaims{
			Audience:  fmt.Sprintf("%v.kohei.com", u.ID),                       // 観測者のUID
			ExpiresAt: time.Now().Add(expiredTime).Unix(),                      // 有効期限14日
			Id:        fmt.Sprintf("%v.kohei.com/%v", u.ID, time.Now().Unix()), // 発行ごとのID
			IssuedAt:  time.Now().Unix(),                                       // 発行日時
			Issuer:    "go.kohei.com",                                          // システムの名前
			NotBefore: time.Now().Unix(),                                       // いつから可能か
			Subject:   fmt.Sprintf("%v.kohei.com", u.ID),                       // 全世界でユニーク
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token.Claims.Valid() != nil {
		return "", token.Claims.Valid()
	}
	jwt, err := token.SignedString(SECRET_KEY)
	return jwt, err
}
