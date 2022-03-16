package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thoriqadillah/e-learning-api/config"
	"github.com/thoriqadillah/e-learning-api/dto"
	"github.com/thoriqadillah/e-learning-api/helpers"
	"github.com/thoriqadillah/e-learning-api/models"
	"github.com/thoriqadillah/e-learning-api/services"
)

type mahasiswaController struct{}

func NewMahasiswaController() services.Crud {
	return &mahasiswaController{}
}

func (m *mahasiswaController) Create(ctx *gin.Context) {
	var mahasiswa models.Mahasiswa

	if err := ctx.ShouldBindJSON(&mahasiswa); err != nil {
		res := []string{}
		for _, er := range err.(validator.ValidationErrors) {
			res = append(res, "Error on field "+er.Field())
		}

		ctx.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(res, nil))
		return
	}

	res := config.DB.Create(&mahasiswa)
	ctx.JSON(http.StatusCreated, helpers.BuildResponse("Success", res.RowsAffected))
}

func (m *mahasiswaController) Read(ctx *gin.Context) {
	var mahasiswa []models.Mahasiswa

	err := config.DB.Find(&mahasiswa).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.BuildResponse("Data not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.BuildResponse("Success", &mahasiswa))
}

func (m *mahasiswaController) ReadOne(ctx *gin.Context) {
	var mahasiswa models.Mahasiswa

	err := config.DB.First(&mahasiswa, ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.BuildResponse("Data not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.BuildResponse("Success", &mahasiswa))
}

func (m *mahasiswaController) Update(ctx *gin.Context) {
	var mahasiswa dto.Mahasiswa
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&mahasiswa); err != nil {
		res := []string{}
		for _, er := range err.(validator.ValidationErrors) {
			res = append(res, "Error on field "+er.Field())
		}

		ctx.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(res, nil))
		return
	}

	err := config.DB.Model(&mahasiswa).Where("id = ?", id).Updates(&mahasiswa).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.BuildResponse("Data not found", nil))
		return
	}

	config.DB.Save(&mahasiswa)
	ctx.JSON(http.StatusOK, helpers.BuildResponse("Success", nil))
}

func (m *mahasiswaController) Delete(ctx *gin.Context) {
	var mahasiswa dto.Mahasiswa
	id := ctx.Param("id")

	err := config.DB.Where("id = ?", id).First(&mahasiswa).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.BuildResponse("Data not found", nil))
		return
	}

	config.DB.Where("id = ?", id).Delete(&mahasiswa)
	ctx.JSON(http.StatusOK, helpers.BuildResponse("Success", nil))
}
