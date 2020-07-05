package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog-go-gin/app/helper"
	"fmt"
)

type Article struct {
	Id          uint32 `db:"id" json:"id"`
	CategoryId  uint8  `db:"category_id" json:"category_id"`
	Title       string `db:"title" json:"title"`
	Author      string `db:"author" json:"author"`
	MarkDown    string `db:"markdown" json:"markdown"`
	Html        string `db:"html" json:"html"`
	Description string `db:"description" json:"description"`
	KeyWords    string `db:"keywords" json:"keywords"`
	Cover       string `db:"cover" json:"cover"`
	IsTop       uint8  `db:"is_top" json:"is_top"`
	IsRecommend uint8  `db:"is_recommend" json:"is_recommend"`
	Views       uint32 `db:"views" json:"views"`
	Status      uint8  `db:"status" json:"status"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
}

func ArticleIndex(c *gin.Context) {
	db := helper.GetDb()
	countSql := "SELECT COUNT(*) as num FROM articles"
	num := 0
	err := db.Get(&num, countSql)
	if err != nil {
		fmt.Printf("count article err:%#v\n", err)
		return
	}
	fmt.Printf("%#v\n", num)

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
