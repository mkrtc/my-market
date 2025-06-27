package retailoutlet

type SeoService struct {
	repo SeoRepoService
}

func NewSeoService(repo SeoRepoService) SeoService {
	return SeoService{repo: repo}
}

func (s *SeoService) FindById(id int) (SeoModel, error) {
	return s.repo.FindById(id)
}

func (s *SeoService) FindAll() ([]SeoModel, error) {
	return s.repo.FindAll()
}

func (s *SeoService) Create(dto CreateSeoDto) (SeoModel, error) {
	model := SeoModel{
		FullName:  dto.FullName,
		ShortName: dto.ShortName,
		OrgName:   dto.OrgName,
	}

	return s.repo.Create(model)
}
