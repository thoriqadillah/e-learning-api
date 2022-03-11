package main

import (
	"github.com/thoriqadillah/e-learning-api/config"
	"github.com/thoriqadillah/e-learning-api/models"
	"gorm.io/gorm"
)

func main() {
	config.DB.AutoMigrate(&models.Mahasiswa{}, &models.Course{})
	seed(config.DB)
}

func seed(db *gorm.DB) {
	mhs := []models.Mahasiswa{
		{NIM: "195150400111034", Name: "Thoriq", Email: "thoriq@student.ub.ac.id", Password: "pass123"},
		{NIM: "185150400111035", Name: "Faisal", Email: "faisal@student.ub.ac.id", Password: "pass123"},
		{NIM: "175150400111036", Name: "Evan", Email: "evan@student.ub.ac.id", Password: "pass123"},
		{NIM: "205150400111037", Name: "Rehan", Email: "rehan@student.ub.ac.id", Password: "pass123"},
		{NIM: "195150400111038", Name: "Nofan", Email: "nofan@student.ub.ac.id", Password: "pass123"},
	}

	db.CreateInBatches(&mhs, len(mhs))
}
