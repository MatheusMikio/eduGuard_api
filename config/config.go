package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = initPostgreSQL()
	if err != nil {
		return fmt.Errorf("error initializing db: %v", err)
	}
	return nil
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}

func GetDb() *gorm.DB {
	return db
}
