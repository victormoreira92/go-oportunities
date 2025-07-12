package handler

import (
	"github.com/victormoreira92/go-oportunities/config"
	"gorm.io/gorm"
)

var (
	looger *config.Logger
	db *gorm.DB
)

func InitializeHandler(){
	looger = config.GetLogger("handler")
	db = config.GetDb()
}
