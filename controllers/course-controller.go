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

//Empty struct for service implementation
type courseController struct{}

func NewCourseController() services.Crud {
	return &courseController{}
}

func (c *courseController) Create(ctx *gin.Context) {
	var course models.Course

	if err := ctx.ShouldBindJSON(&course); err != nil {
		res := []string{}
		for _, er := range err.(validator.ValidationErrors) {
			res = append(res, "Error on field "+er.Field())
		}

		ctx.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(res, nil))
		return
	}

	res := config.DB.Create(&course)
	ctx.JSON(http.StatusOK, helpers.BuildResponse("Success", res.RowsAffected))
}

func (c *courseController) Read(ctx *gin.Context) {
	var course []models.Course

	err := config.DB.Find(&course).Error
	if err != nil {
		panic("Error retrieving course : " + err.Error())
	}

	ctx.JSON(http.StatusOK, helpers.BuildResponse("Success", &course))
}

func (course *courseController) ReadOne(ctx *gin.Context) {}
func (course *courseController) Update(ctx *gin.Context)  {}
func (course *courseController) Delete(ctx *gin.Context)  {}
