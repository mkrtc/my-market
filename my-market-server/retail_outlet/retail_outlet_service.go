package retailoutlet

import (
	"time"
)

type RetailOutletService struct {
	repo RetailOutletRepoService
}

func NewRetailOutletService(repo RetailOutletRepoService) RetailOutletService {
	return RetailOutletService{repo: repo}
}

func (s *RetailOutletService) Create(dto CreateRetailOutletDto) (RetailOutletModel, error) {
	model := RetailOutletModel{
		FullName:   dto.FullName,
		Address:    dto.Address,
		OpenedDate: time.UnixMilli(dto.OpenedDate),
		ClosedDate: time.UnixMilli(dto.ClosedDate),
		SeoId:      int(dto.SeoId),
	}
	return s.repo.Create(model)
}

func (s *RetailOutletService) FindAll() ([]RetailOutletModel, error) {
	return s.repo.FindAll()
}

func (s *RetailOutletService) FindById(id int) (RetailOutletModel, error) {
	return s.repo.FindOne(id)
}
