package user_models

import (
	"errors"
	"fmt"
	"regexp"
)

type Password string

func NewPassword(password *string) (Password, error) {
	if password == nil {
		return "", nil
	}
	// 1. 未入力でない
	if *password == "" {
		return "", errors.New("メールアドレスは必須です。")
	}
	// 2. 254 文字以下である
	const maxLength = 6
	if len(*password) < maxLength {
		return "", fmt.Errorf("パスワードは%d文字以上で入力してください。", maxLength)
	}

	// 3. Password should contain at least one uppercase letter, one lowercase letter, and one digit
	upper := regexp.MustCompile(`[A-Z]`)
	lower := regexp.MustCompile(`[a-z]`)
	digit := regexp.MustCompile(`[0-9]`)
	if !upper.MatchString(*password) || !lower.MatchString(*password) || !digit.MatchString(*password) {
		return "", errors.New("パスワードは少なくとも1つの大文字、1つの小文字、1つの数字を含む必要があります。")
	}

	return Password(*password), nil
}
