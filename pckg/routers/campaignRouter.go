package routers

import (
	campaingPckg "campaing_management/pckg/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func InitializeRouter() {
	router := gin.Default()
	router.Use(CORS())
	campaings := router.Group("campaings")
	campaings.GET("/getCampaingList", campaingPckg.GetCampaings)
	campaings.GET("/getCampaingDetailsById/:id", campaingPckg.GetCampaing)
	campaings.DELETE("/deleteCampaingDetailsById/:id", campaingPckg.DeleteCampaing)
	campaings.POST("/createCampaing", campaingPckg.CreateCampaing)
	campaings.POST("/updateCampaing/:id", campaingPckg.UpdateCampaing)
	campaings.GET("/storeCampaignMockData", campaingPckg.ReadJson)
	router.Run("localhost:8080")
}
