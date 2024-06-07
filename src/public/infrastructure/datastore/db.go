package datastore

import (
	"database/sql"
	"fmt"
	"github.com/kouhei-github/golang-ddd-boboilerplate/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DatabaseProvider interface {
	Connect() (*gorm.DB, *sql.DB, error)
}

type databaseProvider struct {
	URI string
}

func NewDatabaseProvider(env env.Lib) DatabaseProvider {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo",
		env.GetDbUser(),
		env.GetDbPass(),
		env.GetDbHost(),
		env.GetDbPort(),
		env.GetDbName(),
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
