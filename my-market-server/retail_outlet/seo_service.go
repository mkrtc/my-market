package retailoutlet

type SeoService struct {
	repo SeoRepoService
}

func NewSeoService(repo SeoRepoService) *SeoService {
	return &SeoService{repo: repo}
}

func (s *SeoService) FindById(id int) (SeoModel, error) {
	var model SeoModel
	err := s.repo.FindById(id, &model)
	return model, err
}

func (s *SeoService) FindAll() ([]SeoModel, error) {
	var model []SeoModel
	err := s.repo.FindAll(&model)
	return model, err
}

func (s *SeoService) Create(dto CreateSeoDto) (SeoModel, error) {
	model := SeoModel{
		FullName:  dto.FullName,
		ShortName: dto.ShortName,
		OrgName:   dto.OrgName,
	}

	err := s.repo.Create(&model)

	return model, err
}
