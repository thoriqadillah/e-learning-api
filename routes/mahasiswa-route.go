package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqadillah/e-learning-api/controllers"
	"github.com/thoriqadillah/e-learning-api/services"
)

var mahasiswaController services.Crud = controllers.NewMahasiswaController()

func (r *routes) mahasiswaRoute(rg *gin.RouterGroup) {
	r.router.POST("/", mahasiswaController.Create) //add data
	r.router.GET("/", mahasiswaController.Read)    //find all

	mhs := rg.Group("/mahasiswa")
	mhs.GET("/:id", mahasiswaController.Read)      //find one
	mhs.PUT("/:id", mahasiswaController.Update)    //update one
	mhs.DELETE("/:id", mahasiswaController.Delete) //delete one

}
