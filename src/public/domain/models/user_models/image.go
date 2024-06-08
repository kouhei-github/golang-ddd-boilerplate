package user_models

import "errors"

type Image string

func NewImage(image string) (Image, error) {

	if image == "" {
		return "", errors.New("ユーザー画像を正しく入力してください")
	}

	return Image(image), nil
}
