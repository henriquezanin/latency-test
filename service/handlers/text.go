package handlers

import "github.com/gin-gonic/gin"

func (handler Handler) Text(c *gin.Context) {
	t, modelErr := handler.models.Text()
	if modelErr != nil {
		raiseInternalError("failed to handle text request", modelErr)
		writeErrorOnResponse(modelErr, c)
		return
	}
	c.String(200, t)
}
