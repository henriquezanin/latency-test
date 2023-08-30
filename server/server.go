package server

import (
	"TCC2/service"
	"TCC2/utils/logger"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

var production = false

func SetupServer(serviceURL string) *gin.Engine {
	setServiceEnvironment()
	newLogger()
	app := service.New(serviceURL)
	newCors()
	r := newGin()
	loadRoutes(r, app)
	return r
}

func newLogger() {
	logLevel := "debug"
	logFormat := "terminal"
	if production {
		logLevel = "error"
		logFormat = "json"
	}
	logger.SetLevel(logger.Log.ParseLevel(logLevel))
	logger.SetFormat(logFormat)
}

func newGin() *gin.Engine {
	var r *gin.Engine
	if production {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	} else {
		r = gin.Default()
	}
	return r
}

func setServiceEnvironment() {
	prodEnv := os.Getenv("SV_PRODUCTION")
	if prodEnv != "" {
		prod, _ := strconv.ParseBool(prodEnv)
		if prod {
			production = true
		}
	}
}
