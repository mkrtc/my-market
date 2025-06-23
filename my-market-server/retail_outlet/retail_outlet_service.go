package retailoutlet

type RetailOutletService struct {
	service RetailOutletRepoService
}

func NewReatilOutletService(s RetailOutletRepoService) *RetailOutletService {
	return &RetailOutletService{service: s}
}

func (s *RetailOutletService) Create(dto *RetailOutletModel) {
	s.service.Create(dto)
}

func (s *RetailOutletService) FindAll(model *[]RetailOutletModel) {
	s.service.FindAll(model)
}
