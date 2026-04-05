package services

import (
	"github.com/RafehMalik/learning-go-shop/internal/dto"
	"github.com/RafehMalik/learning-go-shop/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetProfile(userID int) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		LastName: user.LastName,
		Phone:    user.Phone,
		Role:     string(user.Role),
		IsActive: user.IsActive,
	}, nil
}
