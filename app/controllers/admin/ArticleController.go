package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/article_list.html", gin.H{})
}

func ArticleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/article_add.html", gin.H{})
}

func ArticleEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/article_edit.html", gin.H{})
}

func ArticleDel(c *gin.Context) {

}