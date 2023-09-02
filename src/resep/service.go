package resep

import (
	"go-resep-api/entity"
	"go-resep-api/src/detailresep"
	"go-resep-api/src/kategori"
)

type Service interface {
	Create(input InputResep) (entity.Resep, error)
	Index() ([]entity.Resep, error)
	FindById(id uint) (entity.Resep, error)
	FindByKategori(kategori string) ([]entity.Resep, error)
	FindByBahan(bahan string) ([]entity.Resep, error)
	Delete(id uint) (bool, error)
	Update(id uint, input InputResep) (entity.Resep, error)
	Restore(id uint) (bool, error)
	PermDelete(id uint) (bool, error)
}

type service struct {
	repository         Repository
	detailRepository   detailresep.Repository
	kategoriRepository kategori.Repository
}

func NewService(repository Repository, detailRepository detailresep.Repository, kategoriRepository kategori.Repository) *service {
	return &service{repository, detailRepository, kategoriRepository}
}

func (s *service) Index() ([]entity.Resep, error) {
	var resep []entity.Resep
	var err error

	resep, err = s.repository.Index()
	if err != nil {
		return resep, err
	}

	return resep, nil
}

func (s *service) FindById(id uint) (entity.Resep, error) {
	var resep entity.Resep
	var err error

	resep, err = s.repository.FindByID(id)

	if err != nil {

		return resep, err
	}

	return resep, nil

}
func (s *service) FindByKategori(kategoriParam string) ([]entity.Resep, error) {
	var resep []entity.Resep
	var kategori entity.Kategori
	var err error

	kategori, err = s.kategoriRepository.FindByNama(kategoriParam)

	if err != nil {
		return resep, err
	}

	resep, err = s.repository.FindByKategory(kategori.ID)

	if err != nil {

		return resep, err
	}

	return resep, nil

}

func (s *service) FindByBahan(bahan string) ([]entity.Resep, error) {
	var resep []entity.Resep
	var err error

	resep, err = s.repository.FindByBahan(bahan)

	if err != nil {

		return resep, err
	}

	return resep, nil

}

func (s *service) Create(input InputResep) (entity.Resep, error) {
	resep := entity.Resep{}
	resep.KategoriId = input.KategoriID
	resep.NamaResep = input.NamaResep

	newresep, err := s.repository.Create(resep)

	if err != nil {
		return newresep, err
	}

	var detailResep []entity.DetailResep

	for _, detail := range input.DetailResep {
		newDetail := entity.DetailResep{
			ResepId: newresep.ID,
			BahanId: detail.BahanID,
			Jumlah:  detail.Jumlah,
		}

		detailResep = append(detailResep, newDetail)
	}

	_, err = s.detailRepository.BatchInsert(detailResep)

	if err != nil {
		return newresep, err
	}

	return newresep, nil
}

func (s *service) Delete(id uint) (bool, error) {
	_, err := s.repository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) Update(id uint, input InputResep) (entity.Resep, error) {
	resep := entity.Resep{}
	resep.ID = id
	resep.NamaResep = input.NamaResep

	updtresep, err := s.repository.Update(resep)

	if err != nil {
		return updtresep, err
	}

	return updtresep, nil
}

func (s *service) Restore(id uint) (bool, error) {
	_, err := s.repository.Restore(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *service) PermDelete(id uint) (bool, error) {
	_, err := s.repository.PermDelete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
