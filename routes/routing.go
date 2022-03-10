package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("api/v1")

	r.authRoute(v1)
	r.mahasiswaRoute(v1)
	r.courseRoute(v1)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}
