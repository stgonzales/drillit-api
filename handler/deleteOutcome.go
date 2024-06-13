package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stgonzles/drillit-api/schemas"
)

func DeleteOutcomeHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errRequiredParam("id", "query parameter").Error())
		return
	}

	outcome := schemas.Outcome{}

	if err := db.First(&outcome, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("outcome id %s not found", id))
		return
	}
	
	if err := db.Delete(&outcome, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting outcome id %s", id))
		return
	}

	sendSuccess(ctx, outcome)
}