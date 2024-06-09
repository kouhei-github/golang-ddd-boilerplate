package external

import "time"

type JWTTokenExternal interface {
	GenerateToken(expiredTime time.Duration, userId int, name, email string) (string, error)
	GetClaimFromToken(token string, secretKey string) (userId int, name, email string, err error)
}
