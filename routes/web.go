package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "world",
		})
	})
}
