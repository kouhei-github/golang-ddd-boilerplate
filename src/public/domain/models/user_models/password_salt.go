package user_models

import "errors"

type PasswordSalt string

func NewPasswordSalt(passwordSalt *string) (PasswordSalt, error) {
	if passwordSalt == nil {
		return "", nil
	}

	if *passwordSalt == "" {
		return "", errors.New("ユーザー画像を正しく入力してください")
	}

	return PasswordSalt(*passwordSalt), nil
}
