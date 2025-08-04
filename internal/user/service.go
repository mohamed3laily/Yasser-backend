package user

import "context"



type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetOrCreateUserByPhone(phone string) (*User, error) {
	user, err := s.repo.FindByPhone(phone)
	if err == nil {
		return user, nil
	}
	return s.repo.Create(phone)
}

func (s *Service) UpdateUser(userID uint, req UpdateUserRequest) (*User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.FullName = req.FullName
	user.FCMToken = req.FCMToken
	user.LanguagePreference = req.LanguagePreference

	updatedUser, err := s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *Service) UpdateLastLogin(ctx context.Context, userID uint) error {
	return s.repo.UpdateLastLogin(ctx, userID)
}

func (s *Service) GetUserByID(userID uint) (*User, error) {
	return s.repo.FindByID(userID)
}
