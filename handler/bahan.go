package handler

import (
	"github.com/gin-gonic/gin"
	"go-resep-api/entity"
	"go-resep-api/helper"
	"go-resep-api/src/bahan"
	"net/http"
)

type bahanHandler struct {
	bahanService bahan.Service
}

func NewBahanHandler(bahanService bahan.Service) *bahanHandler {
	return &bahanHandler{bahanService}
}

func (h *bahanHandler) Index(c *gin.Context) {
	fbahan, err := h.bahanService.Index()

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalahan Saat Mengambil Data", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Mengambil Data", http.StatusOK, "OK", bahan.FormatBahans(fbahan))
	c.JSON(http.StatusOK, response)
	return
}

func (h *bahanHandler) Store(c *gin.Context) {
	var input bahan.InputBahan

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal Input Bahan", http.StatusBadRequest, "Error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newBahan, err := h.bahanService.Create(input)

	if err != nil {
		response := helper.APIResponse("Terjadi kesalah saat menambahkan bahan", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil menambahkan bahan", http.StatusOK, "OK", newBahan)
	c.JSON(http.StatusOK, response)

}

func (h *bahanHandler) Update(c *gin.Context) {
	var inputID bahan.InputBahanID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Mohon masukan id", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusOK, response)

		return
	}

	var fBahan entity.Bahan
	fBahan, err = h.bahanService.FindById(inputID.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan bahan", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fBahan.ID == 0 {
		response := helper.APIResponse("bahan tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	var inputBahan bahan.InputBahan

	err = c.ShouldBindJSON(&inputBahan)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Gagal update bahan", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusOK, response)

		return
	}

	updatedBahan, err := h.bahanService.Update(inputID.ID, inputBahan)

	if err != nil {
		response := helper.APIResponse("Terjadi kesalahan saay update bahan", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil update bahan", http.StatusOK, "OK", updatedBahan)
	c.JSON(http.StatusOK, response)
}

func (h *bahanHandler) Destroy(c *gin.Context) {
	var input bahan.InputBahanID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal Menghapus Bahan", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fBahan entity.Bahan
	fBahan, err = h.bahanService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan bahan", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fBahan.ID == 0 {
		response := helper.APIResponse("bahan tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	isDeleted, err := h.bahanService.Delete(input.ID)

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalah Saat Menghapus Bahan", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Menghapus Bahan", http.StatusOK, "OK", gin.H{"isDeleted": isDeleted})
	c.JSON(http.StatusOK, response)
	return
}

func (h *bahanHandler) PermDestroy(c *gin.Context) {
	var input bahan.InputBahanID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal Menghapus Bahan", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fBahan entity.Bahan
	fBahan, err = h.bahanService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan bahan", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fBahan.ID == 0 {
		response := helper.APIResponse("bahan tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	isDeleted, err := h.bahanService.PermDelete(input.ID)

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalah Saat Menghapus Bahan", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Menghapus Bahan", http.StatusOK, "OK", gin.H{"isDeleted": isDeleted})
	c.JSON(http.StatusOK, response)
	return
}

func (h *bahanHandler) Restore(c *gin.Context) {
	var input bahan.InputBahanID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal Restore Bahan", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fBahan entity.Bahan
	fBahan, err = h.bahanService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan bahan", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fBahan.ID == 0 {
		response := helper.APIResponse("bahan tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	isRestored, err := h.bahanService.Restore(input.ID)

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalah Saat Restore Bahan", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Restore Bahan", http.StatusOK, "OK", gin.H{"isDeleted": isRestored})
	c.JSON(http.StatusOK, response)
	return
}
