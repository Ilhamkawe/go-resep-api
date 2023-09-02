package kategori

import "go-resep-api/entity"

type Service interface {
	Create(input InputKategori) (entity.Kategori, error)
	Index() ([]entity.Kategori, error)
	FindById(id uint) (entity.Kategori, error)
	Delete(id uint) (bool, error)
	Update(id uint, input InputKategori) (entity.Kategori, error)
	Restore(id uint) (bool, error)
	PermDelete(id uint) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]entity.Kategori, error) {
	var kategori []entity.Kategori
	var err error

	kategori, err = s.repository.Index()
	if err != nil {
		return kategori, err
	}

	return kategori, nil
}

func (s *service) FindById(id uint) (entity.Kategori, error) {
	var kategori entity.Kategori
	var err error

	kategori, err = s.repository.FindByID(id)

	if err != nil {

		return kategori, err
	}

	return kategori, nil

}

func (s *service) Create(input InputKategori) (entity.Kategori, error) {
	kategori := entity.Kategori{}
	kategori.Kategori = input.Kategori

	newKategori, err := s.repository.Create(kategori)

	if err != nil {
		return newKategori, err
	}

	return newKategori, nil
}

func (s *service) Delete(id uint) (bool, error) {
	_, err := s.repository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) Update(id uint, input InputKategori) (entity.Kategori, error) {
	kategori := entity.Kategori{}
	kategori.ID = id
	kategori.Kategori = input.Kategori

	updtKategori, err := s.repository.Update(kategori)

	if err != nil {
		return updtKategori, err
	}

	return updtKategori, nil
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
