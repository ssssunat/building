package service

type BuildingService struct {
	repo BuildingRepository
}

func NewService(repo BuildingRepository) *BuildingService {
	return &BuildingService{repo: repo}
}

func (s *BuildingService) CreateBuilding(building Building) (Building, error) {
	return s.repo.CreateBuilding(building)
}

func (s *BuildingService) GetAllBuilding() ([]Building, error) {
	return s.repo.GetAllBuilding()
}