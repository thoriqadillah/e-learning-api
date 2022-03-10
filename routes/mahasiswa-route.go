package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqadillah/e-learning-api/controllers"
	"github.com/thoriqadillah/e-learning-api/services"
)

var mahasiswaController services.Crud = controllers.NewMahasiswaController()

func (r *routes) mahasiswaRoute(rg *gin.RouterGroup) {
	r.router.POST("/", mahasiswaController.Create) //find all
	r.router.GET("/", mahasiswaController.Read)    //add data

	mhs := rg.Group("/mahasiswa")
	mhs.GET("/:id", mahasiswaController.Read)   //find one
	mhs.GET("/:id", mahasiswaController.Update) //update one
	mhs.GET("/:id", mahasiswaController.Delete) //delete one

}
