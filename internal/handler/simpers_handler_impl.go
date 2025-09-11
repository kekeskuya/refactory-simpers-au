package handler

import (
	"errors"
	"fmt"
	"net/http"
	"refactory-simpers-au/config"
	"refactory-simpers-au/constants"
	"refactory-simpers-au/internal/dto"
	"refactory-simpers-au/internal/service"
	"refactory-simpers-au/lib"
	"refactory-simpers-au/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SimpersHandlerImpl struct {
	Env            *config.EnvironmentVariable
	SimpersService service.SimpersService
}

func NewSimpersHandler(
	env *config.EnvironmentVariable,
	simpersService service.SimpersService,
) SimpersHandler {
	return &SimpersHandlerImpl{
		Env:            env,
		SimpersService: simpersService,
	}
}

// GetPersonelByNRP godoc
// @Summary Get Personel berdasarkan NRP
// @Tags Personel
// @Accept json
// @Produce json
// @Param        	nrp             path   			string  	true  "NRP personel"
// @Success 		200 			{object}		lib.APIResponse{data=dto.GetPersonelByNRPResponse}
// @Failure      	400  			{object}  		lib.HTTPError
// @Failure      	500  			{object}  		lib.HTTPError
// @Router /personel/{nrp} [get]
func (h *SimpersHandlerImpl) GetPersonelByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.GetPersonelByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(sqlscan, http.StatusOK, lib.MsgOk, resp)
}

// GetNPWPByNRP godoc
// @Summary Get NPWP berdasarkan NRP
// @Tags Personel
// @Accept json
// @Produce json
// @Param        	nrp             path   			string  	true  "NRP personel"
// @Success 		200 			{object}		lib.APIResponse{data=dto.GetNPWPByNRPResponse}
// @Failure      	400  			{object}  		lib.HTTPError
// @Failure      	500  			{object}  		lib.HTTPError
// @Router /personel/{nrp}/npwp [get]
func (h *SimpersHandlerImpl) GetNPWPByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.GetNPWPByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, constants.DataNPWP), resp)
}

// GetAsabriByNRP godoc
// @Summary Get Asabri berdasarkan NRP
// @Tags Personel
// @Accept json
// @Produce json
// @Param        	nrp             path   			string  	true  "NRP personel"
// @Success 		200 			{object}		lib.APIResponse{data=dto.GetAsabriByNRPResponse}
// @Failure      	400  			{object}  		lib.HTTPError
// @Failure      	500  			{object}  		lib.HTTPError
// @Router /personel/{nrp}/asabri [get]
func (h *SimpersHandlerImpl) GetAsabriByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.GetAsabriByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, constants.DataAsabri), resp)
}

// GetPasporByNRP godoc
// @Summary Get Paspor berdasarkan NRP
// @Tags Personel
// @Accept json
// @Produce json
// @Param        	nrp             path   			string  	true  "NRP personel"
// @Success 		200 			{object}		lib.APIResponse{data=dto.GetPasporByNRPResponse}
// @Failure      	400  			{object}  		lib.HTTPError
// @Failure      	500  			{object}  		lib.HTTPError
// @Router /personel/{nrp}/paspor [get]
func (h *SimpersHandlerImpl) GetPasporByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.GetPasporByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, constants.DataPaspor), resp)
}

// CreateLampiran godoc
// @Summary Menambahkan lampiran berdasarkan id dokumen dan tipe dokumen
// @Tags Lampiran
// @Accept json
// @Produce json
// @Param        	tipe-dokumen    path   			string  					true  "tipe dokumen" Enums(npwp,asabri,paspor)
// @Param        	id-dokumen    	path   			number  					true  "id dokumen"
// @Param         	lampiran        body    		dto.CreateLampiranRequest  	true  "Payload lampiran"
// @Success 		200 			{object}		lib.APIResponse{data=dto.GetPasporByNRPResponse}
// @Failure      	400  			{object}  		lib.HTTPError
// @Failure      	500  			{object}  		lib.HTTPError
// @Router /{tipe-dokumen}/{id-dokumen}/lampiran [post]

func (h *SimpersHandlerImpl) CreateLampiran(sqlscan *gin.Context) {
	docIdStr := sqlscan.Param("id-dokumen")
	if docIdStr == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	docID, err := strconv.Atoi(docIdStr)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	docType := sqlscan.Param("tipe-dokumen")
	if docType == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	if !utils.IsExistsInList(constants.AllowedDocType, docType) {
		err := errors.New("category not supported")
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	var req dto.CreateLampiranRequest
	err = sqlscan.ShouldBindJSON(&req)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.CreateLampiran(req, docID, docType)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgLampiranSuccess, docType), resp)
}
