package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqadillah/e-learning-api/controllers"
	"github.com/thoriqadillah/e-learning-api/services"
)

var mahasiswaController services.Crud = controllers.NewMahasiswaController()

func (r *routes) mahasiswaRoute(rg *gin.RouterGroup) {
	mhs := rg.Group("/mahasiswa")
	mhs.POST("/", mahasiswaController.Create) //add data
	mhs.GET("/", mahasiswaController.Read)    //find all

	oneMhs := rg.Group("/mahasiswa")
	oneMhs.GET("/:id", mahasiswaController.ReadOne)   //find one
	oneMhs.PUT("/:id", mahasiswaController.Update)    //update one
	oneMhs.DELETE("/:id", mahasiswaController.Delete) //delete one
}
