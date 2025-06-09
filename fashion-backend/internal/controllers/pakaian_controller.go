package controllers

import (
	"net/http"
	"strconv"

	"fashion-backend/internal/models"
	"fashion-backend/internal/repositories"
	"fashion-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type PakaianController struct {
	repo repositories.PakaianRepository
}

func NewPakaianController(repo repositories.PakaianRepository) *PakaianController {
	return &PakaianController{repo: repo}
}

func (c *PakaianController) GetAllPakaian(ctx *gin.Context) {
	pakaian, err := c.repo.GetAll()
	if err != nil {
		utils.APIResponse(ctx, "Failed to fetch clothing", http.StatusInternalServerError, nil)
		return
	}
	utils.APIResponse(ctx, "Clothing fetched successfully", http.StatusOK, pakaian)
}

func (c *PakaianController) GetPakaianByID(ctx *gin.Context) {
	id := ctx.Param("id")
	pakaianID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}
	pakaian, err := c.repo.GetByID(uint(pakaianID))
	if err != nil {
		utils.APIResponse(ctx, "Clothing not found", http.StatusNotFound, nil)
		return
	}
	utils.APIResponse(ctx, "Clothing fetched successfully", http.StatusOK, pakaian)
}

func (c *PakaianController) CreatePakaian(ctx *gin.Context) {
	var pakaian models.Pakaian
	if err := ctx.ShouldBindJSON(&pakaian); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	if pakaian.Name == "" || pakaian.Price <= 0 {
		utils.APIResponse(ctx, "Name and valid price are required", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Create(&pakaian); err != nil {
		utils.APIResponse(ctx, "Failed to create clothing", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Clothing created successfully", http.StatusCreated, pakaian)
}

func (c *PakaianController) UpdatePakaian(ctx *gin.Context) {
	id := ctx.Param("id")
	pakaianID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	var pakaian models.Pakaian
	if err := ctx.ShouldBindJSON(&pakaian); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	pakaian.ID = uint(pakaianID)
	if err := c.repo.Update(&pakaian); err != nil {
		utils.APIResponse(ctx, "Failed to update clothing", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Clothing updated successfully", http.StatusOK, pakaian)
}

func (c *PakaianController) DeletePakaian(ctx *gin.Context) {
	id := ctx.Param("id")
	pakaianID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Delete(uint(pakaianID)); err != nil {
		utils.APIResponse(ctx, "Failed to delete clothing", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Clothing deleted successfully", http.StatusOK, nil)
}
