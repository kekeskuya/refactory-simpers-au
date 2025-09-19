package handler

import "github.com/gin-gonic/gin"

type SimpersHandler interface {
	GetPersonelByNRP(sqlscan *gin.Context)
	GetFamilyCardByNRP(sqlscan *gin.Context)
	GetDikMilByNRP(sqlscan *gin.Context)
	GetDikUmByNRP(sqlscan *gin.Context)
	GetJabatanByNRP(sqlscan *gin.Context)
	GetNPWPByNRP(sqlscan *gin.Context)
	GetAsabriByNRP(sqlscan *gin.Context)
	GetPasporByNRP(sqlscan *gin.Context)
	CreateLampiran(sqlscan *gin.Context)
}
