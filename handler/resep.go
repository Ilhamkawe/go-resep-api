package handler

import (
	"github.com/gin-gonic/gin"
	"go-resep-api/entity"
	"go-resep-api/helper"
	"go-resep-api/src/detailresep"
	"go-resep-api/src/resep"
	"net/http"
)

type resepHandler struct {
	resepService  resep.Service
	detailService detailresep.Service
}

func NewResepHandler(resepService resep.Service, detailService detailresep.Service) *resepHandler {
	return &resepHandler{resepService, detailService}
}

func (h *resepHandler) DeleteDetail(c *gin.Context) {
	var input detailresep.InputDetailID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal Menghapus Detail Resep", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fDetail entity.DetailResep
	fDetail, err = h.detailService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan detail", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fDetail.ID == 0 {
		response := helper.APIResponse("detail tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	isDeleted, err := h.detailService.PermDelete(input.ID)

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalah Saat Menghapus detail", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Menghapus detail", http.StatusOK, "OK", gin.H{"isDeleted": isDeleted})
	c.JSON(http.StatusOK, response)
	return
}

func (h *resepHandler) Delete(c *gin.Context) {
	var input resep.InputResepID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal Menghapus Resep", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fDetail entity.Resep
	fDetail, err = h.resepService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan resep", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fDetail.ID == 0 {
		response := helper.APIResponse("resep tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	isDeleted, err := h.resepService.PermDelete(input.ID)

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalah Saat Menghapus resep", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Menghapus resep", http.StatusOK, "OK", gin.H{"isDeleted": isDeleted})
	c.JSON(http.StatusOK, response)
	return
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

	response := helper.APIResponse("Berhasil menemukan resep", http.StatusOK, "OK", resep.FormatResep(fResep))
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
func (h *resepHandler) StoreDetail(c *gin.Context) {
	var input detailresep.InputDetailResep

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal Input detail resep", http.StatusBadRequest, "Error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newResep, err := h.detailService.Create(input)

	if err != nil {
		response := helper.APIResponse("Terjadi kesalah saat menambahkan detail resep", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil menambahkan detail resep", http.StatusOK, "OK", newResep)
	c.JSON(http.StatusOK, response)

}

func (h *resepHandler) Index(c *gin.Context) {
	kategori := c.Query("kategori")
	//bahan := c.Query("bahan")

	if kategori != "" {

		fresep, err := h.resepService.FindByKategori(kategori)

		if err != nil {
			response := helper.APIResponse("Terjadi Kesalahan Saat Mengambil Data", http.StatusInternalServerError, "error", err.Error())
			c.JSON(http.StatusOK, response)
			return
		}

		response := helper.APIResponse("Berhasil Mengambil Data", http.StatusOK, "OK", resep.FormatReseps(fresep))
		c.JSON(http.StatusOK, response)
		return
	}

	//if bahan != "" {
	//	resep, err := h.resepService.FindByBahan(bahan)
	//
	//	if err != nil {
	//		response := helper.APIResponse("Terjadi Kesalahan Saat Mengambil Data", http.StatusInternalServerError, "error", err.Error())
	//		c.JSON(http.StatusOK, response)
	//		return
	//	}
	//
	//	response := helper.APIResponse("Berhasil Mengambil Data", http.StatusOK, "OK", resep)
	//	c.JSON(http.StatusOK, response)
	//	return
	//}

	fresep, err := h.resepService.Index()

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalahan Saat Mengambil Data", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Mengambil Data", http.StatusOK, "OK", resep.FormatReseps(fresep))
	c.JSON(http.StatusOK, response)
	return
}
