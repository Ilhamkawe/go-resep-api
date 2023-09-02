package handler

import (
	"github.com/gin-gonic/gin"
	"go-resep-api/entity"
	"go-resep-api/helper"
	"go-resep-api/src/kategori"
	"net/http"
)

type kategoriHandler struct {
	kategoriService kategori.Service
}

func NewKategoriHandler(kategoriService kategori.Service) *kategoriHandler {
	return &kategoriHandler{kategoriService}
}

func (h *kategoriHandler) Index(c *gin.Context) {
	fkategori, err := h.kategoriService.Index()

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalahan Saat Mengambil Data", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Mengambil Data", http.StatusOK, "OK", kategori.FormatsKategori(fkategori))
	c.JSON(http.StatusOK, response)
	return
}

func (h *kategoriHandler) Store(c *gin.Context) {
	var input kategori.InputKategori

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal Input kategori", http.StatusBadRequest, "Error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newKategori, err := h.kategoriService.Create(input)

	if err != nil {
		response := helper.APIResponse("Terjadi kesalah saat menambahkan kategori", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil menambahkan kategori", http.StatusOK, "OK", newKategori)
	c.JSON(http.StatusOK, response)

}

func (h *kategoriHandler) Update(c *gin.Context) {
	var inputID kategori.InputKategoriID

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

	var fkategori entity.Kategori
	fkategori, err = h.kategoriService.FindById(inputID.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan kategori", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fkategori.ID == 0 {
		response := helper.APIResponse("kategori tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	var inputKategori kategori.InputKategori

	err = c.ShouldBindJSON(&inputKategori)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Gagal update kategori", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusOK, response)

		return
	}

	updatedKategori, err := h.kategoriService.Update(inputID.ID, inputKategori)

	if err != nil {
		response := helper.APIResponse("Terjadi kesalahan saay update kategori", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil update kategori", http.StatusOK, "OK", updatedKategori)
	c.JSON(http.StatusOK, response)
}

func (h *kategoriHandler) Destroy(c *gin.Context) {
	var input kategori.InputKategoriID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal Menghapus kategori", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fKategori entity.Kategori
	fKategori, err = h.kategoriService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan kategori", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fKategori.ID == 0 {
		response := helper.APIResponse("Kategori tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	isDeleted, err := h.kategoriService.Delete(input.ID)

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalah Saat Menghapus Kategori", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Menghapus Kategori", http.StatusOK, "OK", gin.H{"isDeleted": isDeleted})
	c.JSON(http.StatusOK, response)
	return
}

func (h *kategoriHandler) Restore(c *gin.Context) {
	var input kategori.InputKategoriID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal Restore Kategori", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fKategori entity.Kategori
	fKategori, err = h.kategoriService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan bahan", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fKategori.ID == 0 {
		response := helper.APIResponse("kategori tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	isRestored, err := h.kategoriService.Restore(input.ID)

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalah Saat Restore kategori", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Restore kategori", http.StatusOK, "OK", gin.H{"isDeleted": isRestored})
	c.JSON(http.StatusOK, response)
	return
}

func (h *kategoriHandler) PermDestroy(c *gin.Context) {
	var input kategori.InputKategoriID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal Menghapus kategori", http.StatusBadRequest, "Error", errMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	var fKategori entity.Kategori
	fKategori, err = h.kategoriService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Gagal menemukan kategori", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if fKategori.ID == 0 {
		response := helper.APIResponse("kategori tidak ditemukan", http.StatusNotFound, "error", gin.H{})
		c.JSON(http.StatusOK, response)
		return
	}

	isDeleted, err := h.kategoriService.PermDelete(input.ID)

	if err != nil {
		response := helper.APIResponse("Terjadi Kesalah Saat Menghapus Kategori", http.StatusInternalServerError, "Error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Berhasil Menghapus Kategori", http.StatusOK, "OK", gin.H{"isDeleted": isDeleted})
	c.JSON(http.StatusOK, response)
	return
}
