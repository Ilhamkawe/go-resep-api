package handler

import (
	"github.com/gin-gonic/gin"
	"go-resep-api/entity"
	"go-resep-api/helper"
	"go-resep-api/src/resep"
	"net/http"
)

type resepHandler struct {
	resepService resep.Service
}

func NewResepHandler(resepService resep.Service) *resepHandler {
	return &resepHandler{resepService}
}

func (h *resepHandler) FindByID(c *gin.Context) {
	var input resep.InputResepID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal mengambil data", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fResep entity.Resep
	fResep, err = h.resepService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan resep", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil menemukan resep", http.StatusOK, "OK", fResep)
	c.JSON(http.StatusOK, response)

}

func (h *resepHandler) Store(c *gin.Context) {
	var input resep.InputResep

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal Input resep", http.StatusBadRequest, "Error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newResep, err := h.resepService.Create(input)

	if err != nil {
		response := helper.APIResponse("Terjadi kesalah saat menambahkan resep", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil menambahkan resep", http.StatusOK, "OK", newResep)
	c.JSON(http.StatusOK, response)

}

func (h *resepHandler) Index(c *gin.Context) {
	kategori := c.Query("kategori")
	bahan := c.Query("bahan")

	if kategori != "" {

		resep, err := h.resepService.FindByKategori(kategori)

		if err != nil {
			response := helper.APIResponse("Terjadi Kesalahan Saat Mengambil Data", http.StatusInternalServerError, "error", err.Error())
			c.JSON(http.StatusOK, response)
			return
		}

		response := helper.APIResponse("Berhasil Mengambil Data", http.StatusOK, "OK", resep)
		c.JSON(http.StatusOK, response)
		return
	}

	if bahan != "" {
		resep, err := h.resepService.FindByBahan(bahan)

		if err != nil {
			response := helper.APIResponse("Terjadi Kesalahan Saat Mengambil Data", http.StatusInternalServerError, "error", err.Error())
			c.JSON(http.StatusOK, response)
			return
		}

		response := helper.APIResponse("Berhasil Mengambil Data", http.StatusOK, "OK", resep)
		c.JSON(http.StatusOK, response)
		return
	}

	resep, err := h.resepService.Index()

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalahan Saat Mengambil Data", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Mengambil Data", http.StatusOK, "OK", resep)
	c.JSON(http.StatusOK, response)
	return
}
