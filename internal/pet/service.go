package pet

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePet(p Pet) Pet {
	return s.repo.Create(p)
}

func (s *Service) GetAllPets() []Pet {
	return s.repo.GetAll()
}

func (s *Service) GetPetByID(id int) (Pet, bool) {
	return s.repo.GetByID(id)
}
