package handler

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/constants"
	"dummy-simpers-au/internal/service"
	"dummy-simpers-au/lib"
	"fmt"
	"net/http"

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
func (h *SimpersHandlerImpl) GetPersonelByNRP(ctx *gin.Context) {
	nrp := ctx.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.GetPersonelByNRP(nrp)
	if err != nil {
		lib.RespondError(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(ctx, http.StatusOK, lib.MsgOk, resp)
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
func (h *SimpersHandlerImpl) GetNPWPByNRP(ctx *gin.Context) {
	nrp := ctx.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.GetNPWPByNRP(nrp)
	if err != nil {
		lib.RespondError(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(ctx, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, constants.DataNPWP), resp)
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
func (h *SimpersHandlerImpl) GetAsabriByNRP(ctx *gin.Context) {
	nrp := ctx.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.GetAsabriByNRP(nrp)
	if err != nil {
		lib.RespondError(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(ctx, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, constants.DataAsabri), resp)
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
func (h *SimpersHandlerImpl) GetPasporByNRP(ctx *gin.Context) {
	nrp := ctx.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	resp, err := h.SimpersService.GetPasporByNRP(nrp)
	if err != nil {
		lib.RespondError(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(ctx, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, constants.DataPaspor), resp)
}
