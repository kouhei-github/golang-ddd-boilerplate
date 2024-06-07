package user_models

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type User struct {
	ID           interface{}
	UserName     string
	Email        string
	Password     string
	Image        string
	PasswordHash string
	PasswordSalt string
}

// NewUser は、制約に沿った値を持つ仮登録状態のユーザを作成する。
func NewUser(
	email string,
	password string,
	userName string,
	image string,
	passwordHash string,
	passwordSalt string,
) (*User, error) {
	// 識別子である userUUID を生成
	userUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	// ユーザ属性を作成
	return &User{
		ID:           userUUID,
		UserName:     userName,
		Email:        email,
		Password:     password,
		Image:        image,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
	}, nil
}

func (u *User) SetPassword(password string) {
	saltLength := 16
	salt, _ := u.generateSalt(saltLength)
	u.PasswordSalt = salt
	hash, _ := u.hashPassword(password, []byte(salt))
	u.PasswordHash = hash
}

func (u *User) hashPassword(password string, salt []byte) (string, error) {
	pepper := []byte(os.Getenv("PASS_PEPPER"))
	hashedBytes, err := bcrypt.GenerateFromPassword(append(append([]byte(password), salt...), pepper...), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func (u *User) CheckPassword(password string) bool {
	return u.comparePasswords(u.PasswordHash, password, []byte(u.PasswordSalt))
}

func (u *User) comparePasswords(hashedPassword, password string, salt []byte) bool {
	pepper := []byte(os.Getenv("PASS_PEPPER"))
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), append(append([]byte(password), salt...), pepper...))
	return err == nil
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
