package detailresep

import (
	"go-resep-api/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Create(resep entity.DetailResep) (entity.DetailResep, error)
	Index() ([]entity.DetailResep, error)
	Delete(id uint) (bool, error)
	Update(resep entity.DetailResep) (entity.DetailResep, error)
	Restore(id uint) (bool, error)
	FindByID(id uint) (entity.DetailResep, error)
	PermDelete(id uint) (bool, error)
	BatchInsert(resep []entity.DetailResep) ([]entity.DetailResep, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(resep entity.DetailResep) (entity.DetailResep, error) {
	err := r.db.Create(&resep).Error

	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (r *repository) BatchInsert(resep []entity.DetailResep) ([]entity.DetailResep, error) {
	err := r.db.Create(&resep).Error
	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (r *repository) Delete(id uint) (bool, error) {
	err := r.db.Delete(&entity.DetailResep{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) Index() ([]entity.DetailResep, error) {
	var resep []entity.DetailResep

	err := r.db.Order("id asc").Find(&resep).Error

	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (r *repository) Update(resep entity.DetailResep) (entity.DetailResep, error) {
	err := r.db.Save(&resep).Error

	if err != nil {
		return resep, err
	}

	return resep, nil
}

// restore data yang sudah di soft delete
func (r *repository) Restore(id uint) (bool, error) {
	err := r.db.Unscoped().Model(&entity.DetailResep{}).Where("id = ?", id).Update("deleted_at", nil).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) FindByID(id uint) (entity.DetailResep, error) {
	var resep entity.DetailResep

	err := r.db.Unscoped().Where("id = ?", id).Limit(1).Find(&resep).Error
	if err != nil {
		return resep, err
	}

	return resep, nil
}

// function ini akan melakukan permanent delete pada data bahan
func (r *repository) PermDelete(id uint) (bool, error) {
	err := r.db.Unscoped().Delete(&entity.DetailResep{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
