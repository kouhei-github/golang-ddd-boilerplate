package config

import (
	"github.com/joho/godotenv"
)

type ConfigENV interface {
	EnvLoad()
}

type configENV struct {
	dir string
}

func NewConfigENV() ConfigENV {
	dir := "env/"
	return &configENV{
		dir: dir,
	}
}

// EnvLoad は環境変数ファイルを読み込むための関数
func (c configENV) EnvLoad() {
	if err := godotenv.Load(c.dir + ".env"); err != nil {
	} else {
		return
	}
	godotenv.Load(c.dir + ".env.development")
	godotenv.Load(c.dir + ".env.development.local")
}
