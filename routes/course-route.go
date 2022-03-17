package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqadillah/e-learning-api/controllers"
	"github.com/thoriqadillah/e-learning-api/services"
)

var courseController services.Crud = controllers.NewCourseController()

func (r *routes) courseRoute(rg *gin.RouterGroup) {
	course := rg.Group("/course")

	course.POST("/", courseController.Create)      //add data
	course.GET("/", courseController.Read)         //find all
	course.GET("/:id", courseController.ReadOne)   //find one
	course.PUT("/:id", courseController.Update)    //update one
	course.DELETE("/:id", courseController.Delete) //delete one
}
