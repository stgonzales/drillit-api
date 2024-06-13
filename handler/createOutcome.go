package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stgonzles/drillit-api/schemas"
)

func CreateOutcomeHandler(ctx *gin.Context) {
	req := OutcomeRequest{}

	ctx.BindJSON(&req)

	if err := req.ValidateCreateOutcomeRequest(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	outcome := schemas.Outcome{
		Name: req.Name,
		Amount: req.Amount,
	}

	if err := db.Create(&outcome).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error creating outcome")
		return
	}

	sendSuccess(ctx, outcome)
}