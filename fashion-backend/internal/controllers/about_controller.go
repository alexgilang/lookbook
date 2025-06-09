package controllers

import (
	"net/http"
	"strconv"

	"fashion-backend/internal/models"
	"fashion-backend/internal/repositories"
	"fashion-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type AboutController struct {
	repo repositories.AboutRepository
}

func NewAboutController(repo repositories.AboutRepository) *AboutController {
	return &AboutController{repo: repo}
}

func (c *AboutController) GetAbout(ctx *gin.Context) {
	about, err := c.repo.GetAbout()
	if err != nil {
		utils.APIResponse(ctx, "Failed to fetch about information", http.StatusInternalServerError, nil)
		return
	}

	if about == nil {
		utils.APIResponse(ctx, "About information not found", http.StatusNotFound, nil)
		return
	}

	utils.APIResponse(ctx, "About information fetched successfully", http.StatusOK, about)
}

func (c *AboutController) CreateAbout(ctx *gin.Context) {
	var about models.About
	if err := ctx.ShouldBindJSON(&about); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	if about.Name == "" || about.Description == "" || about.PhoneNumber == "" {
		utils.APIResponse(ctx, "Name, description, and phone number are required", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.CreateAbout(&about); err != nil {
		utils.APIResponse(ctx, "Failed to create about information", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "About information created successfully", http.StatusCreated, about)
}

func (c *AboutController) UpdateAbout(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	var about models.About
	if err := ctx.ShouldBindJSON(&about); err != nil {
		utils.APIResponse(ctx, "Invalid request body", http.StatusBadRequest, nil)
		return
	}

	existingAbout, err := c.repo.GetAbout()
	if err != nil {
		utils.APIResponse(ctx, "Failed to fetch about information", http.StatusInternalServerError, nil)
		return
	}

	if existingAbout == nil {
		utils.APIResponse(ctx, "About information not found", http.StatusNotFound, nil)
		return
	}

	existingAbout.Name = about.Name
	existingAbout.Description = about.Description
	existingAbout.PhoneNumber = about.PhoneNumber

	if err := c.repo.UpdateAbout(existingAbout); err != nil {
		utils.APIResponse(ctx, "Failed to update about information", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "About information updated successfully", http.StatusOK, existingAbout)
}

func (c *AboutController) DeleteAbout(ctx *gin.Context) {
	id := ctx.Param("id")
	aboutID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.APIResponse(ctx, "Invalid ID format", http.StatusBadRequest, nil)
		return
	}

	if err := c.repo.DeleteAbout(uint(aboutID)); err != nil {
		utils.APIResponse(ctx, "Failed to delete about information", http.StatusInternalServerError, nil)
		return
	}

	utils.APIResponse(ctx, "About information deleted successfully", http.StatusOK, nil)
}
