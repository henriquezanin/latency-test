package handlers

import "github.com/gin-gonic/gin"

func (handler Handler) Cpu(c *gin.Context) {
	cpu, modelErr := handler.models.Cpu()
	if modelErr != nil {
		raiseInternalError("failed to handle cpu request", modelErr)
		writeErrorOnResponse(modelErr, c)
		return
	}
	c.String(200, cpu)
}
