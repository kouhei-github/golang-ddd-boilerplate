package user_models

import "errors"

type UserName string

func NewUserName(name string) (UserName, error) {

	if name == "" {
		return "", errors.New("ユーザーネームを正しく入力してください")
	}

	return UserName(name), nil
}
