package resep

import (
	"go-resep-api/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Create(resep entity.Resep) (entity.Resep, error)
	Index() ([]entity.Resep, error)
	Delete(id uint) (bool, error)
	Update(resep entity.Resep) (entity.Resep, error)
	Restore(id uint) (bool, error)
	FindByID(id uint) (entity.Resep, error)
	FindByKategory(kategoriID uint) ([]entity.Resep, error)
	PermDelete(id uint) (bool, error)
	FindByBahan(bahan string) ([]entity.Resep, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByName(name string) ([]entity.Resep, error) {
	var resep []entity.Resep

	err := r.db.Where("nama_resep like ?", "%"+name+"%").Find(&resep).Error
	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (r *repository) Create(resep entity.Resep) (entity.Resep, error) {
	err := r.db.Create(&resep).Error

	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (r *repository) Delete(id uint) (bool, error) {
	err := r.db.Delete(&entity.Resep{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) Index() ([]entity.Resep, error) {
	var resep []entity.Resep

	err := r.db.Preload("Kategori").Order("id asc").Find(&resep).Error

	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (r *repository) Update(resep entity.Resep) (entity.Resep, error) {
	err := r.db.Save(&resep).Error

	if err != nil {
		return resep, err
	}

	return resep, nil
}

// restore data yang sudah di soft delete
func (r *repository) Restore(id uint) (bool, error) {
	err := r.db.Unscoped().Model(&entity.Resep{}).Where("id = ?", id).Update("deleted_at", nil).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) FindByBahan(bahan string) ([]entity.Resep, error) {
	var resep []entity.Resep

	err := r.db.Preload("DetailResep", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Bahan").Where("detail_resep.bahan.nama_bahan =?", bahan)
	}).Find(&resep).Error

	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (r *repository) FindByKategory(kategoriID uint) ([]entity.Resep, error) {
	var resep []entity.Resep
	err := r.db.Preload("DetailResep.Bahan").Preload("Kategori").Where("kategori_id = ?", kategoriID).Find(&resep).Error
	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (r *repository) FindByID(id uint) (entity.Resep, error) {
	var resep entity.Resep

	err := r.db.Unscoped().Preload("DetailResep.Bahan").Preload("Kategori").Where("id = ?", id).Limit(1).Find(&resep).Error
	if err != nil {
		return resep, err
	}

	return resep, nil
}

// function ini akan melakukan permanent delete pada data bahan
func (r *repository) PermDelete(id uint) (bool, error) {
	err := r.db.Unscoped().Delete(&entity.Resep{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
