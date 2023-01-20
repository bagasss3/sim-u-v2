package database

import (
	"myapp/config"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgresDB *gorm.DB
)

func InitDB() {
	conn, err := openDB(config.DBDSN())
	if err != nil {
		log.WithField("dbDSN", config.DBDSN()).Fatal("Failed to connect:", err)
	}
	PostgresDB = conn
	log.Info("Success connect database")
}

func openDB(dsn string) (*gorm.DB, error) {
	dialect := postgres.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(config.MaxIdleConns())
	conn.SetMaxOpenConns(config.MaxOpenConns())
	conn.SetConnMaxLifetime(config.ConnMaxLifeTime())
	conn.SetConnMaxIdleTime(config.ConnMaxIdleTime())

	return db, nil
}
