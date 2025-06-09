package controllers

import (
	"net/http"

	"fashion-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (c *HomeController) GetHome(ctx *gin.Context) {
	utils.APIResponse(ctx, "Welcome to Fashion Store API", http.StatusOK, gin.H{
		"message": "This is the home endpoint",
		"routes": []string{
			"GET    /about",
			"GET    /aksesoris",
			"GET    /desain",
			"GET    /pakaian",
			"GET    /parfum",
			"GET    /promosi",
		},
	})
}
