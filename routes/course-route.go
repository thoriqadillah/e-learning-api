package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqadillah/e-learning-api/controllers"
	"github.com/thoriqadillah/e-learning-api/services"
)

var courseController services.Crud = controllers.NewCourseController()

func (r *routes) courseRoute(rg *gin.RouterGroup) {
	course := rg.Group("/course")
	course.POST("/", courseController.Create) //add data
	course.GET("/", courseController.Read)    //find all

	oneCourse := rg.Group("/course")
	oneCourse.GET("/:id", courseController.ReadOne)   //find one
	oneCourse.PUT("/:id", courseController.Update)    //update one
	oneCourse.DELETE("/:id", courseController.Delete) //delete one
}
