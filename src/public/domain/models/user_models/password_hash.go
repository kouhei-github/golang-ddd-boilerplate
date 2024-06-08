package user_models

import "errors"

type PasswordHash string

func NewPasswordHash(passwordHash string) (PasswordHash, error) {

	if passwordHash == "" {
		return "", errors.New("ユーザー画像を正しく入力してください")
	}

	return PasswordHash(passwordHash), nil
}
