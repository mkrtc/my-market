package retailoutlet

import (
	retailoutlet_dto "my-market-server/main/retail_outlet/dto"
	"time"
)

type RetailOutletService struct {
	service RetailOutletRepoService
}

func NewReatilOutletService(s RetailOutletRepoService) *RetailOutletService {
	return &RetailOutletService{service: s}
}

func (s *RetailOutletService) Create(dto retailoutlet_dto.CreateRetailOutletDto) RetailOutletModel {
	model := RetailOutletModel{
		FullName:   dto.FullName,
		Address:    dto.Address,
		OpenedDate: time.UnixMilli(dto.OpenedDate),
		ClosedDate: time.UnixMilli(dto.ClosedDate),
	}
	s.service.Create(&model)

	return model
}

func (s *RetailOutletService) FindAll(model *[]RetailOutletModel) {
	s.service.FindAll(model)
}
