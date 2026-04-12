package server

import (
	"strconv"

	"github.com/RafehMalik/learning-go-shop/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetOrders(c *gin.Context) {
	userID := c.GetUint("user_id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	order, meta, err := s.orderService.GetOrders(userID, page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(c, "no orders found", err)
		return
	}
	utils.PaginatedSuccessResponse(c, "orders retreived successfully", order, *meta)
}

func (s *Server) CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	order, err := s.orderService.CreateOrder(userID)
	if err != nil {
		utils.BadRequestResponse(c, "cannot create order", err)
		return
	}
	utils.CreatedResponse(c, "created successfully", order)
}

func (s *Server) GetOrder(c *gin.Context) {
	userID := c.GetUint("user_id")
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "cannot gt id", err)
		return
	}
	order, er := s.orderService.GetOrder(userID, uint(id))
	if er != nil {
		utils.NotFoundResponse(c, "canot retreive order")
		return
	}
	utils.SuccessResponse(c, "retreived succesfully", order)
}
