package user_models

import (
	"encoding/base64"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
)

type User struct {
	ID           interface{}
	UserName     UserName
	Email        Email
	Password     Password
	Image        Image
	PasswordHash PasswordHash
	PasswordSalt PasswordSalt
}

// NewUser は、制約に沿った値を持つ仮登録状態のユーザを作成する。
func NewUser(
	email Email,
	password Password,
	userName *UserName,
	image *Image,
	passwordHash *PasswordHash,
	passwordSalt *PasswordSalt,
) (*User, error) {
	// 識別子である userUUID を生成
	userUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	user := User{
		ID:       userUUID,
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
		user.PasswordHash = *passwordHash
	}
	if passwordSalt != nil {
		user.PasswordSalt = *passwordSalt
	}

	// ユーザ属性を作成
	return &user, nil
}

func (u *User) GenerateSaltPassword() {
	saltLength := 16
	salt, _ := u.generateSalt(saltLength)
	newSalt, _ := NewPasswordSalt(salt)
	u.PasswordSalt = newSalt
	hash, _ := u.Password.HashPassword([]byte(salt))
	newHash, _ := NewPasswordHash(hash)
	u.PasswordHash = newHash
}

func (u *User) CheckPassword(password string) bool {
	return u.ComparePasswords(string(u.PasswordHash), password, []byte(u.PasswordSalt))
}

func (u *User) generateSalt(length int) (string, error) {
	saltBytes := make([]byte, length)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}

	salt := base64.URLEncoding.EncodeToString(saltBytes)
	return salt, nil
}

func (p *User) ComparePasswords(hashedPassword, password string, salt []byte) bool {
	pepper := []byte(os.Getenv("PASS_PEPPER"))
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), append(append([]byte(password), salt...), pepper...))
	return err == nil
}
