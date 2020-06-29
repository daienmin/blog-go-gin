package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CateIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/cate_list.html", gin.H{})
}

func CateAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/cate_add.html", gin.H{})
}

func CateEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/cate_edit.html", gin.H{})
}

func CateDel(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/cate_add.html", gin.H{})
}