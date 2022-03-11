package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thoriqadillah/e-learning-api/config"
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
		panic("Error retrieving mahasiswa : " + err.Error())
	}

	ctx.JSON(http.StatusOK, helpers.BuildResponse("Success", &mahasiswa))
}

func (m *mahasiswaController) ReadOne(ctx *gin.Context) {}
func (m *mahasiswaController) Update(ctx *gin.Context)  {}
func (m *mahasiswaController) Delete(ctx *gin.Context)  {}
