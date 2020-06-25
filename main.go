package main

import (
	"blog-go-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.WebRoutes(r)
}
