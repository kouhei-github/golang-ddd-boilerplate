package user_models

import (
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"regexp"
)

type Password struct {
	Native string
	Salt   PasswordSalt
	Hash   PasswordHash
}

func NewPassword(password string) (*Password, error) {
	// 1. 未入力でない
	if password == "" {
		return nil, errors.New("パスワードは必須です。")
	}
	// 2. 254 文字以下である
	const maxLength = 6
	if len(password) < maxLength {
		return nil, fmt.Errorf("パスワードは%d文字以上で入力してください。", maxLength)
	}

	// 3. Password should contain at least one uppercase letter, one lowercase letter, and one digit
	upper := regexp.MustCompile(`[A-Z]`)
	lower := regexp.MustCompile(`[a-z]`)
	digit := regexp.MustCompile(`[0-9]`)
	if !upper.MatchString(password) || !lower.MatchString(password) || !digit.MatchString(password) {
		return nil, errors.New("パスワードは少なくとも1つの大文字、1つの小文字、1つの数字を含む必要があります。")
	}

	return &Password{Native: password}, nil
}

func (p *Password) GenerateSaltPassword() {
	saltLength := 16
	salt, _ := p.generateSalt(saltLength)
	newSalt, _ := NewPasswordSalt(salt)
	p.Salt = newSalt
	hash, _ := p.HashPassword([]byte(salt))
	newHash, _ := NewPasswordHash(hash)
	p.Hash = newHash
}

func (p *Password) generateSalt(length int) (string, error) {
	saltBytes := make([]byte, length)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}

	salt := base64.URLEncoding.EncodeToString(saltBytes)
	return salt, nil
}

func (p *Password) HashPassword(salt []byte) (string, error) {
	pepper := []byte(os.Getenv("PASS_PEPPER"))
	hashedBytes, err := bcrypt.GenerateFromPassword(append(append([]byte(p.Native), salt...), pepper...), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}
