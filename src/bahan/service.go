package bahan

import "go-resep-api/entity"

type Service interface {
	Create(input InputBahan) (entity.Bahan, error)
	Delete(id uint) (bool, error)
	PermDelete(id uint) (bool, error)
	Restore(id uint) (bool, error)
	Index() ([]entity.Bahan, error)
	Update(id uint, input InputBahan) (entity.Bahan, error)
	FindById(id uint) (entity.Bahan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]entity.Bahan, error) {
	var bahan []entity.Bahan
	var err error

	bahan, err = s.repository.Index()

	if bahan != nil {
		return bahan, err
	}

	return bahan, nil
}

func (s *service) Create(input InputBahan) (entity.Bahan, error) {
	bahan := entity.Bahan{}
	bahan.NamaBahan = input.NamaBahan
	bahan.Satuan = input.Satuan

	newBahan, err := s.repository.Create(bahan)

	if err != nil {
		return newBahan, err
	}

	return newBahan, nil
}

func (s *service) Update(id uint, input InputBahan) (entity.Bahan, error) {

	bahan := entity.Bahan{}
	bahan.ID = id
	bahan.NamaBahan = input.NamaBahan
	bahan.Satuan = input.Satuan

	updtBahan, err := s.repository.Update(bahan)
	if err != nil {
		return updtBahan, err
	}

	return updtBahan, nil
}

func (s *service) FindById(id uint) (entity.Bahan, error) {
	var bahan entity.Bahan
	var err error

	bahan, err = s.repository.FindByID(id)

	if err != nil {

		return bahan, err
	}

	return bahan, nil

}

func (s *service) Delete(id uint) (bool, error) {
	_, err := s.repository.Delete(id)
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
func (s *service) Restore(id uint) (bool, error) {
	_, err := s.repository.Restore(id)
	if err != nil {
		return false, err
	}
	return true, nil
}
