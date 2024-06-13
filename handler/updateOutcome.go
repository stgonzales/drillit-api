package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stgonzles/drillit-api/schemas"
)

func UpdateOutcomeHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errRequiredParam("id", "query parameter").Error())
		return
	}

	req := OutcomeRequest{}

	ctx.BindJSON(&req)

	if err := req.ValidateUpdateOutcomeRequest(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	outcome := schemas.Outcome{}

	if err := db.First(&outcome, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("outcome id %s not found", id))
		return
	}

	if req.Name != "" {
		outcome.Name = req.Name
	}

	if req.Amount > 0 {
		outcome.Amount = req.Amount
	}

	if err := db.Save(&outcome).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("outcome id %s not found", id))
		return
	}

	sendSuccess(ctx, outcome)
}