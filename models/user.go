package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"fname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
}
