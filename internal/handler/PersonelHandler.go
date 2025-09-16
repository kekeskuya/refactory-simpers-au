package handler

import (
	"net/http"
	"refactory-simpers-au/internal/dto"
	"refactory-simpers-au/internal/repository"

	"github.com/gin-gonic/gin"
)

type PersonelHandler struct {
	repo repository.PersonelRepository
}

func NewPersonelHandler(repo repository.PersonelRepository) *PersonelHandler {
	return &PersonelHandler{repo: repo}
}

// GET /personel/:id
func (h *PersonelHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	personel, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	resp := dto.ToGetPersonelByNRPResponse(personel)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "OK",
		"data":    resp,
	})
}

// GET /personel/nrp/:nrp
func (h *PersonelHandler) GetByNRP(c *gin.Context) {
	nrp := c.Param("nrp")

	personel, err := h.repo.GetByNRP(nrp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	resp := dto.ToGetPersonelByNRPResponse(personel)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "OK",
		"data":    resp,
	})
}
