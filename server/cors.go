package server

import (
	"github.com/gin-contrib/cors"
	"log"
	"os"
)

var corsConfig cors.Config

func newCors() {
	corsConfig = cors.DefaultConfig()
	//if production {
	//	corsProduction()
	//} else {
	//	corsDevelopment()
	//}
	corsDevelopment()
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"POST", "GET", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Accept", "Content-Type", "Content-Length", "Authorization"}
}

func corsProduction() {
	url := os.Getenv("SV_DOMAIN")
	if url == "" {
		log.Fatal("null env SV_DOMAIN")
	}
	corsConfig.AllowOrigins = []string{url}
}

func corsDevelopment() {
	corsConfig.AllowOrigins = []string{"*"}
}
