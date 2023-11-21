package models

import (
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	Name string
}
