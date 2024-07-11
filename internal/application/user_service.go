package application

import "Azgart/internal/domain"

type UserService struct {
}

func (s *UserService) GetUserByID(userID string) *domain.User {
	return &domain.User{ID: userID, Name: "John Doe"}
}
