package user_models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Email string

func NewEmail(email string) (Email, error) {
	// 1. 未入力でない
	if email == "" {
		return "", errors.New("メールアドレスは必須です。")
	}
	// 2. 254 文字以下である
	const maxLength = 254
	if len(email) > maxLength {
		return "", fmt.Errorf("メールアドレスは%d文字以下で入力してください。", maxLength)
	}
	// 3. @が1つだけ含まれている
	if strings.Count(email, "@") != 1 {
		return "", errors.New("メールアドレスには「@」が１つだけ含まれている必要があります。")
	}
	// @前後で分割する
	splitEmail := strings.Split(email, "@")
	localPart := splitEmail[0]
	domainPart := splitEmail[1]
	// 4. ローカル部は1文字以上64文字以下である
	const (
		minimumLength = 1
		maximumLength = 64
	)
	if utf8.RuneCountInString(localPart) < minimumLength || maximumLength < utf8.RuneCountInString(localPart) {
		return "", fmt.Errorf("「@」の前は%d文字以上%d文字以下で入力してください。", minimumLength, maximumLength)
	}
	// 5. ローカル部が英数字とピリオドのみから構成される
	r := regexp.MustCompile(`^[\w.]+$`)
	if !r.MatchString(localPart) {
		return "", errors.New("「@」の前は半角英数字、ピリオドのみで入力してください。")
	}
	// 6. ローカル部の先頭がピリオドで始まっていない
	if strings.HasPrefix(localPart, ".") {
		return "", errors.New("「@」の前の先頭にピリオドは使用できません。")
	}
	// 7. ローカル部の末尾がピリオドで終わっていない
	if strings.HasSuffix(localPart, ".") {
		return "", errors.New("「@」の前の末尾にピリオドは使用できません。")
	}
	// 8. ローカル部でピリオドが連続していない
	if strings.Contains(localPart, "..") {
		return "", errors.New("「@」の前でピリオドは連続して使用できません。")
	}
	// 9. ドメイン部がドメインの形式として正しい
	r = regexp.MustCompile(`^[a-zA-Z0-9-]{1,63}(\.[a-zA-Z0-9-]{1,63})*\.[a-zA-Z]{2,}$`)
	if !r.MatchString(domainPart) {
		return "", errors.New("「@」の後はドメインの形式で入力してください。")
	}
	return Email(email), nil
}

func (e Email) UserName() string {
	emails := strings.Split(string(e), "@")
	return emails[0]
}
