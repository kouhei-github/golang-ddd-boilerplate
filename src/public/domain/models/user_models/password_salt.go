package user_models

import "errors"

type PasswordSalt string

func NewPasswordSalt(passwordSalt string) (PasswordSalt, error) {

	if passwordSalt == "" {
		return "", errors.New("ソルトキーを正しく入力してください")
	}

	return PasswordSalt(passwordSalt), nil
}
