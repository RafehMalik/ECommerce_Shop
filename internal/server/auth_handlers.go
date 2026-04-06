package server

import (
	"github.com/RafehMalik/learning-go-shop/internal/dto"
	"github.com/RafehMalik/learning-go-shop/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	response, err := s.authService.Register(&req)
	if err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	utils.CreatedResponse(c, "user creates successfully", response)
}

func (s *Server) login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	response, err := s.authService.Login(&req)
	if err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	utils.CreatedResponse(c, "login successfull", response)
}

func (s *Server) refreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	response, err := s.authService.RefreshToken(&req)
	if err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	utils.CreatedResponse(c, "login successfull", response)
}

func (s *Server) logout(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	err := s.authService.Logout(req.RefreshToken)
	if err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	utils.CreatedResponse(c, "logout successfull", nil)
}

func (s *Server) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	profile, err := s.userService.GetProfile(userID)
	if err != nil {
		utils.NotFoundResponse(c, "record not found")
		return
	}
	utils.SuccessResponse(c, "retreived succesfully", profile)
}

func (s *Server) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid data", err)
		return
	}

	profile, err := s.userService.UpdateProfle(userID, &req)
	if err != nil {
		utils.BadRequestResponse(c, "failed to update data", err)
		return
	}
	utils.SuccessResponse(c, "updatd successfuly", profile)
}
