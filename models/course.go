package models

import (
	"gorm.io/gorm"
)

//Course database table reperesentation for creation
type Course struct {
	gorm.Model
	Kode      string       `json:"kode" binding:"required"`
	Name      string       `json:"name" binding:"required"`
	Mahasiswa []*Mahasiswa `gorm:"many2many:mahasiswa_course"`
}
