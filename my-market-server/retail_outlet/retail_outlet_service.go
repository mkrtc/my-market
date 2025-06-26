package retailoutlet

import (
	"time"
)

type RetailOutletService struct {
	repo RetailOutletRepoService
}

func NewRetailOutletService(repo RetailOutletRepoService) *RetailOutletService {
	return &RetailOutletService{repo: repo}
}

func (s *RetailOutletService) Create(dto CreateRetailOutletDto) (RetailOutletModel, error) {
	model := RetailOutletModel{
		FullName:   dto.FullName,
		Address:    dto.Address,
		OpenedDate: time.UnixMilli(dto.OpenedDate),
		ClosedDate: time.UnixMilli(dto.ClosedDate),
		SeoId:      int(dto.SeoId),
	}
	err := s.repo.Create(&model)

	return model, err
}

func (s *RetailOutletService) FindAll() ([]RetailOutletModel, error) {
	var model []RetailOutletModel
	err := s.repo.FindAll(&model)
	return model, err
}

func (s *RetailOutletService) FindById(id int) (RetailOutletModel, error) {
	var model RetailOutletModel
	err := s.repo.FindOne(id, &model)
	return model, err
}
