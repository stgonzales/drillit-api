package config

import (
	"os"

	"github.com/stgonzles/drillit-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	logger.Info("GORM connecting to db and automigrating")


	dbPath := "./db/sqlite.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("DB file not found, creating...")
		
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Failed to create folder 'db': %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create DB file: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to open SQLite file: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Outcome{})
	if err != nil {
		logger.Errorf("Failed to run automigration: %v", err)
		return nil, err
	}

	logger.Info("GORM connection and automigration completed!")
	return db, nil
}