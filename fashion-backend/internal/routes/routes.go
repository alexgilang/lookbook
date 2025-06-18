package routes

import (
	"fashion-backend/internal/controllers"
	"fashion-backend/internal/repositories"
	"fashion-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		utils.APIResponse(ctx, "Welcome to Fashion API", http.StatusOK, nil)
	})

	aboutRepo := repositories.NewAboutRepository(db)
	aksesorisRepo := repositories.NewAksesorisRepository(db)
	desainRepo := repositories.NewDesainRepository(db)
	pakaianRepo := repositories.NewPakaianRepository(db)
	parfumRepo := repositories.NewParfumRepository(db)
	promosiRepo := repositories.NewPromosiRepository(db)

	aboutController := controllers.NewAboutController(aboutRepo)
	aksesorisController := controllers.NewAksesorisController(aksesorisRepo)
	desainController := controllers.NewDesainController(desainRepo)
	pakaianController := controllers.NewPakaianController(pakaianRepo)
	parfumController := controllers.NewParfumController(parfumRepo)
	promosiController := controllers.NewPromosiController(promosiRepo)

	aboutGroup := router.Group("/about")
	{
		aboutGroup.GET("/", aboutController.GetAbout)
		aboutGroup.POST("/", aboutController.CreateAbout)
		aboutGroup.PUT("/:id", aboutController.UpdateAbout)
		aboutGroup.DELETE("/:id", aboutController.DeleteAbout)
	}

	aksesorisGroup := router.Group("/aksesoris")
	{
		aksesorisGroup.GET("/", aksesorisController.GetAllAksesoris)
		aksesorisGroup.POST("/", aksesorisController.CreateAksesoris)
		aksesorisGroup.GET("/:id", aksesorisController.GetAksesorisByID)
		aksesorisGroup.PUT("/:id", aksesorisController.UpdateAksesoris)
		aksesorisGroup.DELETE("/:id", aksesorisController.DeleteAksesoris)
	}

	desainGroup := router.Group("/desain")
	{
		desainGroup.GET("/", desainController.GetAllDesain)
		desainGroup.POST("/", desainController.CreateDesain)
		desainGroup.GET("/:id", desainController.GetDesainByID)
		desainGroup.PUT("/:id", desainController.UpdateDesain)
		desainGroup.DELETE("/:id", desainController.DeleteDesain)
	}

	pakaianGroup := router.Group("/pakaian")
	{
		pakaianGroup.GET("/", pakaianController.GetAllPakaian)
		pakaianGroup.POST("/", pakaianController.CreatePakaian)
		pakaianGroup.GET("/:id", pakaianController.GetPakaianByID)
		pakaianGroup.PUT("/:id", pakaianController.UpdatePakaian)
		pakaianGroup.DELETE("/:id", pakaianController.DeletePakaian)
	}

	parfumGroup := router.Group("/parfum")
	{
		parfumGroup.GET("/", parfumController.GetAllParfum)
		parfumGroup.POST("/", parfumController.CreateParfum)
		parfumGroup.GET("/:id", parfumController.GetParfumByID)
		parfumGroup.PUT("/:id", parfumController.UpdateParfum)
		parfumGroup.DELETE("/:id", parfumController.DeleteParfum)
	}

	promosiGroup := router.Group("/promosi")
	{
		promosiGroup.GET("/", promosiController.GetAllPromosi)
		promosiGroup.POST("/", promosiController.CreatePromosi)
		promosiGroup.GET("/:id", promosiController.GetPromosiByID)
		promosiGroup.PUT("/:id", promosiController.UpdatePromosi)
		promosiGroup.DELETE("/:id", promosiController.DeletePromosi)
	}
}
