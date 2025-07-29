package handler

import "github.com/gin-gonic/gin"

type SimpersHandler interface {
	GetPersonelByNRP(ctx *gin.Context)
	GetNPWPByNRP(ctx *gin.Context)
	GetAsabriByNRP(ctx *gin.Context)
	GetPasporByNRP(ctx *gin.Context)
	CreateLampiran(ctx *gin.Context)
}
