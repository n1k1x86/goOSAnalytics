package main

import (
	"osAnalytics/pkg/infoapi"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/static", "./templates/static")

	infoapi.RegisterRoutes(r)

	r.Run(":8000")
}
