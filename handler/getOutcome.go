package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOutcomesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "outcome",
	})
}