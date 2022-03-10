package models

import (
	"gorm.io/gorm"
)

//Course database table reperesentation
type Course struct {
	gorm.Model
	Kode        string `json:"kode" binding:"required"`
	Name        string `json:"name" binding:"required"`
	MahasiswaID uint   `json:"mahasiswa_id,omitempty" binding:"required"`
	Mahasiswa   Mahasiswa
}
