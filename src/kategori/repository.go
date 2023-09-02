package kategori

import (
	"go-resep-api/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Create(kategori entity.Kategori) (entity.Kategori, error)
	Index() ([]entity.Kategori, error)
	Delete(id uint) (bool, error)
	Update(kategori entity.Kategori) (entity.Kategori, error)
	Restore(id uint) (bool, error)
	FindByID(id uint) (entity.Kategori, error)
	FindByNama(nama string) (entity.Kategori, error)
	PermDelete(id uint) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(kategori entity.Kategori) (entity.Kategori, error) {
	err := r.db.Create(&kategori).Error

	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (r *repository) Delete(id uint) (bool, error) {
	err := r.db.Delete(&entity.Kategori{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) Index() ([]entity.Kategori, error) {
	var kategori []entity.Kategori

	err := r.db.Order("id asc").Find(&kategori).Error

	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (r *repository) Update(kategori entity.Kategori) (entity.Kategori, error) {
	err := r.db.Save(&kategori).Error

	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

// restore data yang sudah di soft delete
func (r *repository) Restore(id uint) (bool, error) {
	err := r.db.Unscoped().Model(&entity.Kategori{}).Where("id = ?", id).Update("deleted_at", nil).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) FindByNama(nama string) (entity.Kategori, error) {
	var kategori entity.Kategori

	err := r.db.Where("kategori = ?", nama).Limit(1).Find(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (r *repository) FindByID(id uint) (entity.Kategori, error) {
	var kategori entity.Kategori

	err := r.db.Unscoped().Where("id = ?", id).Limit(1).Find(&kategori).Error
	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

// function ini akan melakukan permanent delete pada data bahan
func (r *repository) PermDelete(id uint) (bool, error) {
	err := r.db.Unscoped().Delete(&entity.Kategori{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
