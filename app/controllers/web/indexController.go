package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home1/index.html", gin.H{})
}

func Category(c *gin.Context) {
	c.HTML(http.StatusOK, "home1/list.html", gin.H{})
}

func Article(c *gin.Context) {
	c.HTML(http.StatusOK, "home1/info.html", gin.H{})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "home/about.html", gin.H{})
}

func Share(c *gin.Context) {
	c.HTML(http.StatusOK, "home/share.html", gin.H{})
}

func ShareInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "home/infopic.html", gin.H{})
}