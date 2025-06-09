package controllers

import (
	"net/http"
	"strconv"

	"fashion-backend/internal/models"
	"fashion-backend/internal/repositories"
	"fashion-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type ParfumController struct {
	repo repositories.ParfumRepository
}

func NewParfumController(repo repositories.ParfumRepository) *ParfumController {
	return &ParfumController{repo: repo}
}

func (c *ParfumController) GetAllParfum(ctx *gin.Context) {
	parfums, err := c.repo.GetAll()
	if err != nil {
		utils.APIResponse(ctx, "Failed to fetch perfumes", http.StatusInternalServerError, nil)
		return
	}
	utils.APIResponse(ctx, "Perfumes fetched successfully", http.StatusOK, parfums)
}

func (c *ParfumController) GetParfumByID(ctx *gin.Context) {
	id := ctx.Param("id")
	parfumID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}
	parfum, err := c.repo.GetByID(uint(parfumID))
	if err != nil {
		utils.APIResponse(ctx, "Perfume not found", http.StatusNotFound, nil)
		return
	}
	utils.APIResponse(ctx, "Perfume fetched successfully", http.StatusOK, parfum)
}

func (c *ParfumController) CreateParfum(ctx *gin.Context) {
	var parfum models.Parfum
	if err := ctx.ShouldBindJSON(&parfum); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	if parfum.Name == "" || parfum.Price <= 0 {
		utils.APIResponse(ctx, "Name and valid price are required", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Create(&parfum); err != nil {
		utils.APIResponse(ctx, "Failed to create perfume", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Perfume created successfully", http.StatusCreated, parfum)
}

func (c *ParfumController) UpdateParfum(ctx *gin.Context) {
	id := ctx.Param("id")
	parfumID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	var parfum models.Parfum
	if err := ctx.ShouldBindJSON(&parfum); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	parfum.ID = uint(parfumID)
	if err := c.repo.Update(&parfum); err != nil {
		utils.APIResponse(ctx, "Failed to update perfume", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Perfume updated successfully", http.StatusOK, parfum)
}

func (c *ParfumController) DeleteParfum(ctx *gin.Context) {
	id := ctx.Param("id")
	parfumID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Delete(uint(parfumID)); err != nil {
		utils.APIResponse(ctx, "Failed to delete perfume", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Perfume deleted successfully", http.StatusOK, nil)
}
