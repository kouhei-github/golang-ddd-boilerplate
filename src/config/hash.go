package config

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"hash/fnv"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var PEPPER = []byte(os.Getenv("PASS_PEPPER"))

// Sha256 (64桁)
func GetBinaryBySHA256WithKey(msg string, key string) ([]byte, error) {
	mac := hmac.New(sha256.New, getBinaryBySHA256(key))
	_, err := mac.Write([]byte(msg))
	return mac.Sum(nil), err
}

func getBinaryBySHA256(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}

// fnv (8桁)
func Hash32ByFnv(message []byte) string {
	h := fnv.New32a()
	h.Write(message)
	sum := h.Sum32()
	result := strconv.FormatUint(uint64(sum), 16)
	return result
}

func ParsePassword(password string) string {
	hashedPass := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(hashedPass[:])
}

func GenerateSalt(length int) (string, error) {
	saltBytes := make([]byte, length)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}

	salt := base64.URLEncoding.EncodeToString(saltBytes)
	return salt, nil
}

func HashPassword(password string, salt []byte) (string, error) {
	pepper := PEPPER
	hashedBytes, err := bcrypt.GenerateFromPassword(append(append([]byte(password), salt...), pepper...), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func ComparePasswords(hashedPassword, password string, salt []byte) bool {
	pepper := PEPPER
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), append(append([]byte(password), salt...), pepper...))
	return err == nil
}
