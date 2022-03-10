package models

import (
	"gorm.io/gorm"
)

//Mahasiswa database table reperesentation
type Mahasiswa struct {
	gorm.Model
	NIM      string `json:"nim" binding:"required,min=15,max=15"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,max=15,alphanum" `
}
