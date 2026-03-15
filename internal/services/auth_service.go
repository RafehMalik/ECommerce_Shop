package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/RafehMalik/learning-go-shop/internal/config"
	"github.com/RafehMalik/learning-go-shop/internal/dto"
	"github.com/RafehMalik/learning-go-shop/internal/models"
	"github.com/RafehMalik/learning-go-shop/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	db     *gorm.DB
	config *config.Config
}

func NewAuthService(db *gorm.DB, confi *config.Config) *AuthService {
	return &AuthService{
		db:     db,
		config: confi,
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	var existingUser models.User
	err := s.db.Where("email=?", req.Email).First(existingUser).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user := models.User{
		Email:     req.Email,
		Password:  hashPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Role:      models.UserRoleCustomer,
	}
	err = s.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	cart := models.Cart{
		UserID: user.ID,
	}
	err = s.db.Create(cart).Error
	if err != nil {
		fmt.Println("unalbe to create cart")
	}

	return s.generateAuthResponse(&user)
}

func (s *AuthService) login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	var user models.User
	err := s.db.Where("email=? and is_active=?", req.Email, true).First(&user).Error
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("password incorect")
	}
	return s.generateAuthResponse(&user)
}

func (s *AuthService) RefreshToken(req *dto.RefreshTokenRequest) (*dto.AuthResponse, error) {
	claims, err := utils.ValidateToken(req.RefreshToken, s.config.JWT.Secret)
	if err != nil {
		return nil, err
	}
	var refreshToken models.RefreshToken
	err = s.db.Where("token=? And expires_at=?", req.RefreshToken, time.Now()).First(&refreshToken).Error
	if err != nil {
		return nil, errors.New("toen not exst")
	}
	var user models.User
	err = s.db.First(&user, claims.UserID).Error
	if err != nil {
		return nil, errors.New("user not exst")
	}
	s.db.Delete(&refreshToken)
	return s.generateAuthResponse(&user)
}

func (s *AuthService) logout(refreshToken string) error {

	return s.db.Where("token=?", refreshToken).Delete(&models.RefreshToken{}).Error
}
func (s *AuthService) generateAuthResponse(user *models.User) (*dto.AuthResponse, error) {
	accessToken, refreshToen, err := utils.GenerateTokenPair(
		&s.config.JWT,
		user.ID,
		user.Email,
		string(user.Role),
	)
	if err != nil {
		return nil, err
	}
	refreshToenmodel := models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToen,
		ExpiresAt: time.Now().Add(s.config.JWT.RefreshTokenExpires),
	}
	s.db.Create(refreshToenmodel)
	return &dto.AuthResponse{
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			Role:      string(user.Role),
			IsActive:  user.IsActive,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToen,
	}, nil

}
