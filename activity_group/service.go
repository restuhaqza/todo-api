package activity_group

type Service interface {
	GetAll() ([]ActivityGroup, error)
	GetByID(ID string) (ActivityGroup, error)
	Create(activityGroup ActivityGroup) (ActivityGroup, error)
	Delete(ID string) (bool, error)
	Update(activityGroup ActivityGroup) (ActivityGroup, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]ActivityGroup, error) {
	return s.repository.FindAll()
}

func (s *service) GetByID(ID string) (ActivityGroup, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(activityGroup ActivityGroup) (ActivityGroup, error) {
	return s.repository.Create(activityGroup)
}

func (s *service) Delete(ID string) (bool, error) {
	return s.repository.Delete(ID)
}

func (s *service) Update(activityGroup ActivityGroup) (ActivityGroup, error) {
	return s.repository.Update(activityGroup)
}
