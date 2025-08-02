package user

type Service interface {
	GetOrCreateUserByPhone(phone string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetOrCreateUserByPhone(phone string) (*User, error) {
	user, err := s.repo.FindByPhone(phone)
	if err == nil {
		return user, nil
	}
	return s.repo.Create(phone)
}
