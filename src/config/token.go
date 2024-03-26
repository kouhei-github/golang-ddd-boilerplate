package config

import (
	"os"
	"time"
)

var unit = time.Minute

var (
	JWT_SECRET_KEY      = []byte(os.Getenv("JWT_SECRET_KEY"))
	RefreshTokenExpires = 60 * 24 * 30 * 2 * unit
	AccessTokenExpires  = 20 * unit
)
