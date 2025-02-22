package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stgonzles/drillit-api/handler"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	v1 := api.Group("/v1")
	{
		v1.GET("/list-outcomes", handler.ListOutcomesHandler)
		v1.GET("/get-outcome", handler.GetOutcomeHandler)
		v1.POST("/create-outcome", handler.CreateOutcomeHandler)
		v1.PATCH("/update-outcome", handler.UpdateOutcomeHandler)
		v1.DELETE("/delete-outcome", handler.DeleteOutcomeHandler)
	}
}