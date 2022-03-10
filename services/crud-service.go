package services

import "github.com/gin-gonic/gin"

type Crud interface {
	Create(ctx *gin.Context)
	Read(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
