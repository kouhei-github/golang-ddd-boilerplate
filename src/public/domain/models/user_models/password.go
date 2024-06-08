package user_models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
)

type Password string

func NewPassword(password string) (Password, error) {
	// 1. 未入力でない
	if password == "" {
		return "", errors.New("パスワードは必須です。")
	}
	// 2. 254 文字以下である
	const maxLength = 6
	if len(password) < maxLength {
		return "", fmt.Errorf("パスワードは%d文字以上で入力してください。", maxLength)
	}

	// 3. Password should contain at least one uppercase letter, one lowercase letter, and one digit
	upper := regexp.MustCompile(`[A-Z]`)
	lower := regexp.MustCompile(`[a-z]`)
	digit := regexp.MustCompile(`[0-9]`)
	if !upper.MatchString(password) || !lower.MatchString(password) || !digit.MatchString(password) {
		return "", errors.New("パスワードは少なくとも1つの大文字、1つの小文字、1つの数字を含む必要があります。")
	}

	return Password(password), nil
}

func (p *Password) HashPassword(salt []byte) (string, error) {
	pepper := []byte(os.Getenv("PASS_PEPPER"))
	hashedBytes, err := bcrypt.GenerateFromPassword(append(append([]byte(*p), salt...), pepper...), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}
