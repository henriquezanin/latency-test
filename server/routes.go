package server

import (
	"TCC2/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func loadRoutes(r *gin.Engine, app service.Application) {
	if production {
		r.Use(jsonLoggerMiddleware())
	}
	r.Use(cors.New(corsConfig))
	r.GET("/cpu", app.Routes.Cpu)
	r.GET("/text", app.Routes.Text)
}
