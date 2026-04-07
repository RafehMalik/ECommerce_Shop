package server

import (
	"strconv"

	"github.com/RafehMalik/learning-go-shop/internal/dto"
	"github.com/RafehMalik/learning-go-shop/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	category, err := s.productService.CreateCateory(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to create", err)
		return
	}
	utils.SuccessResponse(c, "created successfully", category)
}

func (s *Server) GetCategory(c *gin.Context) {
	category, err := s.productService.GetCategory()
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to get", err)
		return
	}
	utils.SuccessResponse(c, "retrieved succefully", category)
}

func (s *Server) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid category id", err)
		return
	}
	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	category, err := s.productService.UpdateCategory(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to update", err)
		return
	}
	utils.SuccessResponse(c, "succefully updated", category)
}

func (s *Server) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid category id", err)
		return
	}
	err = s.productService.DeleteCategory(int(id))
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to delete", err)
		return
	}
	utils.SuccessResponse(c, "deleted succesfully", nil)
}

func (s *Server) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "invalid request", err)
		return
	}
	product, err := s.productService.CreateProduct(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to create", err)
		return
	}
	utils.SuccessResponse(c, "created successfully", product)
}

func (s *Server) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "invalid product id", err)
		return
	}
	category, err := s.productService.GetProduct(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "product not found")
		return
	}
	utils.SuccessResponse(c, "retrieved succefully", category)
}

func (s *Server) GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	category, meta, err := s.productService.GetProducts(page, limit)
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to get", err)
		return
	}
	utils.PaginatedSuccessResponse(c, "retrieved succefully", category, *meta)
}

func (s *Server) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID", err)
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	product, err := s.productService.UpdateProduct(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update product", err)
		return
	}
	utils.SuccessResponse(c, "Product updated successfully", product)
}

func (s *Server) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID", err)
		return
	}
	if err := s.productService.DeleteProduct(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete product", err)
		return
	}

	utils.SuccessResponse(c, "Product deleted successfully", nil)
}

func (s *Server) uploadProductImage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID", err)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		utils.BadRequestResponse(c, "No file uploaded", err)
		return
	}

	url, err := s.uploadService.UploadProductImage(uint(id), file)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to upload image", err)
		return
	}

	if err := s.productService.AddProductImage(uint(id), url, file.Filename); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to save image record", err)
		return
	}

	utils.SuccessResponse(c, "Image uploaded successfully", map[string]string{"url": url})
}
