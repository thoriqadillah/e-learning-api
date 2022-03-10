package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqadillah/e-learning-api/controllers"
	"github.com/thoriqadillah/e-learning-api/services"
)

var courseController services.Crud = controllers.NewCourseController()

func (r *routes) courseRoute(rg *gin.RouterGroup) {
	r.router.POST("/", courseController.Create) //add data
	r.router.GET("/", courseController.Read)    //find all

	course := rg.Group("/course")
	course.GET("/:id", courseController.Read)      //find one
	course.PUT("/:id", courseController.Update)    //update one
	course.DELETE("/:id", courseController.Delete) //delete one
}
