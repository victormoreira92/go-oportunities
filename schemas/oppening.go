package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Oppening struct {
	gorm.Model
	Role string
	Company string
	Location string
	Remote bool
	Link string 
	Salary int64
}

type OppeningResponse struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt time.Time `json:"updateAt" `
	Role 	string `json:"role"`
	Company 	string `json:"company" `
	Location 	string `json:"location" `
	Remote 	bool `json:"remote"`
	Link 	string  `json:"link"`
	Salary 	int64 `json:"salary"`
	DeletedAt time.Time `json:deletedAt,omitempty` 
}