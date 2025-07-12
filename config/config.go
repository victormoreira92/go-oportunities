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
	db, err = InitializeSqlLite()
	if err != nil{
		return fmt.Errorf("error initializing db: %v", err)
	}
	return nil 
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger{
	logger = NewLogger(p)
	return logger
}