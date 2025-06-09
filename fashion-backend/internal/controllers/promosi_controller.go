package controllers

import (
	"net/http"
	"strconv"

	"fashion-backend/internal/models"
	"fashion-backend/internal/repositories"
	"fashion-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type PromosiController struct {
	repo repositories.PromosiRepository
}

func NewPromosiController(repo repositories.PromosiRepository) *PromosiController {
	return &PromosiController{repo: repo}
}

func (c *PromosiController) GetAllPromosi(ctx *gin.Context) {
	promos, err := c.repo.GetAll()
	if err != nil {
		utils.APIResponse(ctx, "Failed to fetch promotions", http.StatusInternalServerError, nil)
		return
	}
	utils.APIResponse(ctx, "Promotions fetched successfully", http.StatusOK, promos)
}

func (c *PromosiController) GetPromosiByID(ctx *gin.Context) {
	id := ctx.Param("id")
	promoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	promo, err := c.repo.GetByID(uint(promoID))
	if err != nil {
		utils.APIResponse(ctx, "Promotion not found", http.StatusNotFound, nil)
		return
	}
	utils.APIResponse(ctx, "Promotion fetched successfully", http.StatusOK, promo)
}

func (c *PromosiController) CreatePromosi(ctx *gin.Context) {
	var promosi models.Promosi
	if err := ctx.ShouldBindJSON(&promosi); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	if promosi.Name == "" || promosi.Price <= 0 || promosi.StartDate.IsZero() || promosi.EndDate.IsZero() {
		utils.APIResponse(ctx, "Name, price and valid start and end dates are required", http.StatusBadRequest, nil)
		return
	}

	if promosi.EndDate.Before(promosi.StartDate) {
		utils.APIResponse(ctx, "End date must be after start date", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.Create(&promosi); err != nil {
		utils.APIResponse(ctx, "Failed to create promotion", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Promotion created successfully", http.StatusCreated, promosi)
}

func (c *PromosiController) UpdatePromosi(ctx *gin.Context) {
	id := ctx.Param("id")
	promoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	var promosi models.Promosi
	if err := ctx.ShouldBindJSON(&promosi); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	if !promosi.EndDate.IsZero() && !promosi.StartDate.IsZero() && promosi.EndDate.Before(promosi.StartDate) {
		utils.APIResponse(ctx, "End date must be after start date", http.StatusBadRequest, nil)
		return
	}

	promosi.ID = uint(promoID)
	if err := c.repo.Update(&promosi); err != nil {
		utils.APIResponse(ctx, "Failed to update promotion", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "Promotion updated successfully", http.StatusOK, promosi)
}

func (c *PromosiController) DeletePromosi(ctx *gin.Context) {
	id := ctx.Param("id")
	promoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}
	if err := c.repo.Delete(uint(promoID)); err != nil {
		utils.APIResponse(ctx, "Failed to delete promotion", http.StatusInternalServerError, nil)
		return
	}
	utils.APIResponse(ctx, "Promotion deleted successfully", http.StatusOK, nil)
}
