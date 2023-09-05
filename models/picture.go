package models

import (
	"gorm.io/gorm"
)

type Picture struct {
	gorm.Model
	Src string `json:"src"` 
	UserID uint
  }  