package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Class string `json:"class"`
}
