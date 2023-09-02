package bahan

import (
	"go-resep-api/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Create(bahan entity.Bahan) (entity.Bahan, error)
	Delete(id uint) (bool, error)
	PermDelete(id uint) (bool, error)
	Restore(id uint) (bool, error)
	Index() ([]entity.Bahan, error)
	FindByID(id uint) (entity.Bahan, error)
	Update(bahan entity.Bahan) (entity.Bahan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(bahan entity.Bahan) (entity.Bahan, error) {

	err := r.db.Create(&bahan).Error
	if err != nil {
		return bahan, err
	}

	return bahan, err
}

// function ini akan melakukan soft delete pada data bahan
func (r *repository) Delete(id uint) (bool, error) {
	err := r.db.Delete(&entity.Bahan{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// restore data yang sudah di soft delete
func (r *repository) Restore(id uint) (bool, error) {
	err := r.db.Unscoped().Model(&entity.Bahan{}).Where("id = ?", id).Update("deleted_at", nil).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// function ini akan melakukan permanent delete pada data bahan
func (r *repository) PermDelete(id uint) (bool, error) {
	err := r.db.Unscoped().Delete(&entity.Bahan{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) IndexPaginated(limit int, offset int) ([]entity.Bahan, error) {
	var bahan []entity.Bahan

	err := r.db.Limit(limit).Offset(offset).Order("id asc").Find(&bahan).Error
	if err != nil {
		return bahan, err
	}

	return bahan, nil

}

func (r *repository) Index() ([]entity.Bahan, error) {
	var bahan []entity.Bahan

	err := r.db.Order("id asc").Find(&bahan).Error
	if err != nil {
		return bahan, err
	}

	return bahan, nil
}

func (r *repository) FindByID(id uint) (entity.Bahan, error) {
	var bahan entity.Bahan

	err := r.db.Unscoped().Where("id = ?", id).Limit(1).Find(&bahan).Error
	if err != nil {
		return bahan, err
	}

	return bahan, nil
}

func (r *repository) Update(bahan entity.Bahan) (entity.Bahan, error) {
	err := r.db.Save(&bahan).Error
	if err != nil {
		return bahan, err
	}
	return bahan, nil
}
