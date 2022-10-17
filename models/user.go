package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreatedAt time.Time `json:"cAt,omitempty"`
	FirstName string    `json:"fname,"`
	LastName  string    `json:"lastname,omitempty"`
}
