package controllers

import (
	"net/http"
	"strconv"

	"fashion-backend/internal/models"
	"fashion-backend/internal/repositories"
	"fashion-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type AksesorisController struct {
	repo repositories.AksesorisRepository
}

func NewAksesorisController(repo repositories.AksesorisRepository) *AksesorisController {
	return &AksesorisController{repo: repo}
}

func (c *AksesorisController) GetAllAksesoris(ctx *gin.Context) {
	aksesoris, err := c.repo.GetAll()
	if err != nil {
		utils.APIResponse(ctx, "Failed to fetch accessories", http.StatusInternalServerError, nil)
		return
	}
	utils.APIResponse(ctx, "Accessories fetched successfully", http.StatusOK, aksesoris)
}

func (c *AksesorisController) GetAksesorisByID(ctx *gin.Context) {
	id := ctx.Param("id")
	aksesorisID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}
	aksesoris, err := c.repo.GetByID(uint(aksesorisID))
	if err != nil {
		utils.APIResponse(ctx, "Accessory not found", http.StatusNotFound, nil)
		return
	}
	utils.APIResponse(ctx, "Accessory fetched successfully", http.StatusOK, aksesoris)
}

func (c *AksesorisController) CreateAksesoris(ctx *gin.Context) {
	var aksesoris models.Aksesoris
	if err := ctx.ShouldBindJSON(&aksesoris); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	if aksesoris.Name == "" || aksesoris.Price <= 0 {
		utils.APIResponse(ctx, "Name and valid price are required", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Create(&aksesoris); err != nil {
		utils.APIResponse(ctx, "Failed to create accessory", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Accessory created successfully", http.StatusCreated, aksesoris)
}

func (c *AksesorisController) UpdateAksesoris(ctx *gin.Context) {
	id := ctx.Param("id")
	aksesorisID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	var aksesoris models.Aksesoris
	if err := ctx.ShouldBindJSON(&aksesoris); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	aksesoris.ID = uint(aksesorisID)
	if err := c.repo.Update(&aksesoris); err != nil {
		utils.APIResponse(ctx, "Failed to update accessory", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Accessory updated successfully", http.StatusOK, aksesoris)
}

func (c *AksesorisController) DeleteAksesoris(ctx *gin.Context) {
	id := ctx.Param("id")
	aksesorisID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Delete(uint(aksesorisID)); err != nil {
		utils.APIResponse(ctx, "Failed to delete accessory", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Accessory deleted successfully", http.StatusOK, nil)
}
