package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("Failed to initialize SQLite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLgger(p)
	return logger
}