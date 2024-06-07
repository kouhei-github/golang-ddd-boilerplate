package env

import (
	"fmt"
	"os"
)

type Lib interface {
	CheckValues() error
	GetDbRoot() string
	GetDbName() string
	GetDbUser() string
	GetDbPass() string
	GetDbHost() string
	GetDbPort() string
	GetJwtSecret() string
}

type lib struct {
	DbRoot    string
	DbName    string
	DbUser    string
	DbPass    string
	DbHost    string
	DbPort    string
	JwtSecret string
}

func NewLib() Lib {
	return lib{
		DbRoot:    os.Getenv("DB_ROOT"),
		DbName:    os.Getenv("DB_NAME"),
		DbUser:    os.Getenv("DB_USER"),
		DbPass:    os.Getenv("DB_PASS"),
		DbHost:    os.Getenv("DB_HOST"),
		DbPort:    os.Getenv("DB_PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func (l lib) CheckValues() error {
	if l.DbRoot == "" {
		fmt.Println("DB_ROOT is not set")
		return fmt.Errorf("DB_ROOT is not set")
	}
	if l.DbName == "" {
		fmt.Println("DB_NAME is not set")
		return fmt.Errorf("DB_NAME is not set")
	}
	if l.DbUser == "" {
		fmt.Println("DB_USER is not set")
		return fmt.Errorf("DB_USER is not set")
	}
	if l.DbPass == "" {
		fmt.Println("DB_PASS is not set")
		return fmt.Errorf("DB_PASS is not set")
	}
	if l.DbHost == "" {
		fmt.Println("DB_HOST is not set")
		return fmt.Errorf("DB_HOST is not set")
	}
	if l.DbPort == "" {
		fmt.Println("DB_PORT is not set")
		return fmt.Errorf("DB_PORT is not set")
	}
	if l.JwtSecret == "" {
		fmt.Println("JWT_SECRET is not set")
		return fmt.Errorf("JWT_SECRET is not set")
	}
	return nil
}

func (l lib) GetDbRoot() string {
	return l.DbRoot
}

func (l lib) GetDbName() string {
	return l.DbName
}

func (l lib) GetDbUser() string {
	return l.DbUser
}

func (l lib) GetDbPass() string {
	return l.DbPass
}

func (l lib) GetDbHost() string {
	return l.DbHost
}

func (l lib) GetDbPort() string {
	return l.DbPort
}

func (l lib) GetJwtSecret() string {
	return l.JwtSecret
}
