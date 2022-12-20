package routers

import (
	campaingPckg "campaing_management/pckg/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	router := gin.Default()
	campaings := router.Group("campaings")
	campaings.GET("/getCampaings", campaingPckg.GetCampaings)
	campaings.GET("/:id", campaingPckg.GetCampaing)
	campaings.DELETE("/:id", campaingPckg.DeleteCampaing)
	campaings.POST("/createCampaing", campaingPckg.CreateCampaing)
	campaings.POST("/update/:id", campaingPckg.UpdateCampaing)
	campaings.GET("/storeMockData", campaingPckg.ReadJson)
	router.Run("localhost:8080")
}
