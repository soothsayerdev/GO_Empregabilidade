package router

import (
	"github.com/gin-gonic/gin"
	"github.com/soothsayerdev/GO_Empregabilidade/handler"
)

// InitializeRoutes initializes the routes for the application
//						pointer of router in gin library
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1") // Declare version of api
	{
		v1.GET("/opening", handler.ShowOpeningHandler)  
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler) 		
		v1.GET("/openings", handler.ListOpeningHandler)  
	
	}
}