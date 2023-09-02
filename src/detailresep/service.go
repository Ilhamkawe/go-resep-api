package detailresep

import "go-resep-api/entity"

type Service interface {
	PermDelete(id uint) (bool, error)
	FindById(id uint) (entity.DetailResep, error)
	Create(input InputDetailResep) (entity.DetailResep, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input InputDetailResep) (entity.DetailResep, error) {
	detail := entity.DetailResep{}
	detail.BahanId = input.BahanID
	detail.ResepId = input.ResepID
	detail.Jumlah = input.Jumlah

	newDetail, err := s.repository.Create(detail)

	if err != nil {
		return newDetail, err
	}

	return newDetail, nil
}

func (s *service) PermDelete(id uint) (bool, error) {
	_, err := s.repository.PermDelete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) FindById(id uint) (entity.DetailResep, error) {
	var detail entity.DetailResep
	var err error

	detail, err = s.repository.FindByID(id)

	if err != nil {

		return detail, err
	}

	return detail, nil

}
