package config

import (
	"os"

	"github.com/victormoreira92/go-oportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlLite() (*gorm.DB, error){
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_,err := os.Stat(dbPath)

	if os.IsNotExist(err){
		logger.Info("database file not found created....")
		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil{
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil{
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath),&gorm.Config{})
	
	if err != nil{
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Oppening{})
	if err != nil{
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	return db, nil
}