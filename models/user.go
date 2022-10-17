package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"fname,"`
	LastName  string `json:"lastname,omitempty"`
}
