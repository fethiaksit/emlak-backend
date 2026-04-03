package routes

import (
	"emlak-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ListingRoutes(router *gin.Engine) {
	api := router.Group("/api/listings")
	{
		api.GET("", controllers.GetListings)
		api.GET("/:id", controllers.GetListingByID)
		api.POST("", controllers.CreateListing)
		api.DELETE("/:id", controllers.DeleteListing)
	}
}
