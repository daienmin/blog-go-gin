package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LinkIndex(c *gin.Context)  {
	c.HTML(http.StatusOK, "admin/link_index.html", gin.H{})
}

func LinkAdd(c *gin.Context)  {
	c.HTML(http.StatusOK, "admin/link_add.html", gin.H{})
}

func LinkCreate(c *gin.Context)  {}

func LinkEdit(c *gin.Context)  {
	c.HTML(http.StatusOK, "admin/link_edit.html", gin.H{})
}

func LinkUpdate(c *gin.Context)  {}

func LinkDel(c *gin.Context)  {}