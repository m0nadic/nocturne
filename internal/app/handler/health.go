package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func (hh HealthHandler) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
