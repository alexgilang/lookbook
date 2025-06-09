package controllers

import (
	"net/http"
	"strconv"

	"fashion-backend/internal/models"
	"fashion-backend/internal/repositories"
	"fashion-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type DesainController struct {
	repo repositories.DesainRepository
}

func NewDesainController(repo repositories.DesainRepository) *DesainController {
	return &DesainController{repo: repo}
}

func (c *DesainController) CreateDesain(ctx *gin.Context) {
	var desain models.Desain
	if err := ctx.ShouldBindJSON(&desain); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Create(&desain); err != nil {
		utils.APIResponse(ctx, "Failed to create Desain", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Desain created successfully", http.StatusCreated, desain)
}

func (c *DesainController) GetAllDesain(ctx *gin.Context) {
	desains, err := c.repo.GetAll()
	if err != nil {
		utils.APIResponse(ctx, "Failed to fetch Desains", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Desains fetched successfully", http.StatusOK, desains)
}

func (c *DesainController) GetDesainByID(ctx *gin.Context) {
	id := ctx.Param("id")
	desainID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	desain, err := c.repo.GetByID(uint(desainID))
	if err != nil {
		utils.APIResponse(ctx, "Desain not found", http.StatusNotFound, nil)
		return
	}

	utils.APIResponse(ctx, "Desain fetched successfully", http.StatusOK, desain)
}

func (c *DesainController) UpdateDesain(ctx *gin.Context) {
	id := ctx.Param("id")
	desainID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	var desain models.Desain
	if err := ctx.ShouldBindJSON(&desain); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	desain.ID = uint(desainID)
	if err := c.repo.Update(&desain); err != nil {
		utils.APIResponse(ctx, "Failed to update Desain", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Desain updated successfully", http.StatusOK, desain)
}

func (c *DesainController) DeleteDesain(ctx *gin.Context) {
	id := ctx.Param("id")
	desainID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Delete(uint(desainID)); err != nil {
		utils.APIResponse(ctx, "Failed to delete Desain", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Desain deleted successfully", http.StatusOK, nil)
}
