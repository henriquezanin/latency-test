package handlers

import (
	"TCC2/service/handlers/models"
	"TCC2/utils/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	models Models
}

type Models interface {
	Cpu() (string, *models.Error)
}

func New(models Models) Handler {
	return Handler{models: models}
}

func raiseInternalError(message string, err *models.Error) {
	if err.InternalErr == nil && err.InternalErrLevel < logger.Fatal {
		return
	}
	event := logger.Log.NewEvent(err.InternalErrLevel)
	if err.InternalErrLevel <= logger.Error {
		event.AddField("error", err.InternalErr)
	} else {
		event.AddField("info", err.InternalErr)
	}
	event.Log(message)
}

func writeErrorOnResponse(err *models.Error, c *gin.Context) {
	if err.HttpErr == "" && err.HttpStatusCode == 0 {
		return
	}
	if err.HttpErr == "" && err.HttpStatusCode != 0 {
		c.Status(err.HttpStatusCode)
		return
	}
	c.JSON(err.HttpStatusCode, gin.H{"message": err.HttpErr})
}
