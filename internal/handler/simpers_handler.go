package handler

import "github.com/gin-gonic/gin"

type SimpersHandler interface {
	GetPersonelByNRP(ctx *gin.Context)
	GetNPWPByNRP(ctx *gin.Context)
}
