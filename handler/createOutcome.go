package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stgonzles/drillit-api/schemas"
)

func CreateOutcomeHandler(ctx *gin.Context) {
	req := CreateOutcomeRequest{}

	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	outcome := schemas.Outcome{
		Name: req.Name,
		Amount: req.Amount,
	}

	if err := db.Create(&outcome).Error; err != nil {
		logger.Errorf("Failed to create outcome: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating outcome")
		return
	}

	sendSuccess(ctx, outcome)
}