package database

import (
	"fmt"
	"futbook/app/model"
	"futbook/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB holds the database
type DB struct{ *gorm.DB }

// database instance
var defaultDB = &DB{}

// connect sets the db client of database using configuration
func (db *DB) connect(cfg *config.DB) (err error) {
	dbUri := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Name,
	)
	db.DB, err = gorm.Open(mysql.Open(dbUri), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	// Try to ping database.
	if err != nil {
		return fmt.Errorf("can't sent ping to database, %w", err)
	}

	return nil
}

// GetDB returns db instance
func GetDB() *DB {
	return defaultDB
}

// ConnectDB sets the db client of database using default configuration
func ConnectDB() error {
	return defaultDB.connect(config.DBCfg())
}
