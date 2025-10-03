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

// GetLampiranKKByNRP implements SimpersHandler.

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
	fmt.Printf("DEBUG NRP: %s\n", nrp)
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

func (h *SimpersHandlerImpl) GetFamilyCardByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}
	resp, err := h.SimpersService.GetFamilyCardByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}
	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Kartu Keluarga"), resp)
}

func (h *SimpersHandlerImpl) GetDikMilByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}
	resp, err := h.SimpersService.GetDikMilByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}
	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Pendidikan Militer"), resp)
}

func (h *SimpersHandlerImpl) GetDikUmByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}
	resp, err := h.SimpersService.GetDikUmByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}
	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Pendidikan Umum"), resp)
}

func (h *SimpersHandlerImpl) GetJabatanByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}
	resp, err := h.SimpersService.GetJabatanByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}
	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Jabatan"), resp)
}

func (h *SimpersHandlerImpl) GetTanhorByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}
	resp, err := h.SimpersService.GetTanhorByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}
	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Tanda Honor"), resp)
}

func (h *SimpersHandlerImpl) GetPangkatByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}
	resp, err := h.SimpersService.GetPangkatByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}
	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Pangkat"), resp)
}

func (h *SimpersHandlerImpl) GetLampiranKKByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}
	resp, err := h.SimpersService.GetLampiranKKByNRP(nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}
	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Lampiran KK"), resp)
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

func (h *SimpersHandlerImpl) GetDokumenByNRP(sqlscan *gin.Context) {
	nrp := sqlscan.Param("nrp")
	doc_type := sqlscan.Param("tipe_dokumen")

	if nrp == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	if doc_type == "" {
		err := constants.ErrorMessageInvalidInput
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	// resp, err := h.SimpersService.GetDokumenByNRP(nrp)
	// if err != nil {
	// 	lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
	// 	return
	// }

	switch doc_type {
	case "kartu_keluarga":
		//fmt.Println("DEBUG DOC TYPE :", doc_type)
		resp, err := h.SimpersService.GetFamilyCardByNRP(nrp)
		if err != nil {
			lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
			return
		}
		lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Kartu Keluarga"), resp)
	case "pendidikan_militer":
		//fmt.Println("DEBUG DOC TYPE :", doc_type)
		resp, err := h.SimpersService.GetDikMilByNRP(nrp)
		if err != nil {
			lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
			return
		}
		lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Pendidikan Militer"), resp)
	case "pendidikan_umum":
		//fmt.Println("DEBUG DOC TYPE :", doc_type)
		resp, err := h.SimpersService.GetDikUmByNRP(nrp)
		if err != nil {
			lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
			return
		}
		lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Pendidikan Umum"), resp)
	case "jabatan":
		//fmt.Println("DEBUG DOC TYPE :", doc_type)
		resp, err := h.SimpersService.GetJabatanByNRP(nrp)
		if err != nil {
			lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
			return
		}
		lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Jabatan"), resp)
	case "tanhor":
		//fmt.Println("DEBUG DOC TYPE :", doc_type)
		resp, err := h.SimpersService.GetTanhorByNRP(nrp)
		if err != nil {
			lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
			return
		}
		lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Tanda Honor"), resp)
	case "pangkat":
		//fmt.Println("DEBUG DOC TYPE :", doc_type)
		resp, err := h.SimpersService.GetPangkatByNRP(nrp)
		if err != nil {
			lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
			return
		}
		lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgDokumenSuccess, "Pangkat"), resp)
	default:
		fmt.Println("DEBUG DOC TYPE :")
		err := errors.New("document category not supported")
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return

	}

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

	resp, err := h.SimpersService.CreateLampiran(req, strconv.Itoa(docID), docType)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgLampiranSuccess, docType), resp)
}

// CreateLampiranDokumen implements SimpersHandler interface
// func (h *SimpersHandlerImpl) CreateLampiranDokumen(ctx *gin.Context) {
// 	// TODO: Implement the actual logic or delegate to an existing method if appropriate
// 	lib.RespondError(ctx, http.StatusNotImplemented, "not implemented", nil)
// }

func (h *SimpersHandlerImpl) CreateLampiranDokumen(sqlscan *gin.Context) {
	var req dto.DokumenLampiranByNRPRequest
	if err := sqlscan.ShouldBindJSON(&req); err != nil {
		lib.RespondError(sqlscan, http.StatusBadRequest, err.Error(), err)
		return
	}

	nrp := sqlscan.Param("nrp")

	resp, err := h.SimpersService.CreateLampiranDokumen(req, nrp)
	if err != nil {
		lib.RespondError(sqlscan, http.StatusInternalServerError, err.Error(), err)
		return
	}

	lib.RespondSuccess(sqlscan, http.StatusOK, fmt.Sprintf(lib.MsgLampiranSuccess, "Data Lampiran"), resp)
}
