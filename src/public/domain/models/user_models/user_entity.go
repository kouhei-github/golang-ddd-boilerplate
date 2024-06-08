package user_models

import (
	"golang.org/x/crypto/bcrypt"
	"os"
)

type User struct {
	ID       int
	UserName UserName
	Email    Email
	Password Password
	Image    Image
}

// NewUser は、制約に沿った値を持つ仮登録状態のユーザを作成する。
func NewUser(
	id int,
	email Email,
	password Password,
	userName *UserName,
	image *Image,
	passwordHash *PasswordHash,
	passwordSalt *PasswordSalt,
) (*User, error) {

	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if userName != nil {
		user.UserName = *userName
	}
	if image != nil {
		user.Image = *image
	}
	if passwordHash != nil {
		user.Password.Hash = *passwordHash
	}
	if passwordSalt != nil {
		user.Password.Salt = *passwordSalt
	}

	// ユーザ属性を作成
	return &user, nil
}

func (u *User) CheckPassword(password string) bool {
	return u.ComparePasswords(string(u.Password.Hash), password, []byte(u.Password.Salt))
}

func (p *User) ComparePasswords(hashedPassword, password string, salt []byte) bool {
	pepper := []byte(os.Getenv("PASS_PEPPER"))
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), append(append([]byte(password), salt...), pepper...))
	return err == nil
}
