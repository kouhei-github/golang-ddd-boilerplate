package providers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseProvider interface {
	Connect() (*gorm.DB, *sql.DB, error)
}

type databaseProvider struct {
	URI string
}

func NewDatabaseProvider() DatabaseProvider {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return &databaseProvider{
		URI: uri,
	}
}

func (p databaseProvider) Connect() (*gorm.DB, *sql.DB, error) {
	conn, err := p.sqlConnect()
	if err != nil {
		log.Printf("err at Connect: %s", err)
		return nil, nil, err
	}
	db, err := conn.DB()
	if err != nil {
		return nil, nil, err
	}

	return conn, db, nil
}

func (p databaseProvider) sqlConnect() (database *gorm.DB, err error) {
	return gorm.Open(mysql.Open(p.URI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
