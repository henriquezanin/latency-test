package handlers

import "github.com/gin-gonic/gin"

func (handler Handler) Cpu(c *gin.Context) {
	cpu, modelErr := handler.models.Cpu()
	if modelErr != nil {
		raiseInternalError("user login has failed", modelErr)
		writeErrorOnResponse(modelErr, c)
		return
	}
	c.String(200, cpu)
}
