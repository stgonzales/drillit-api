package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stgonzles/drillit-api/schemas"
)

func GetOutcomesHandler(ctx *gin.Context) {
	outcomes := []schemas.Outcome{}

	if err := db.Find(&outcomes).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing outcomes")
		return
	}

	sendSuccess(ctx, outcomes)
}