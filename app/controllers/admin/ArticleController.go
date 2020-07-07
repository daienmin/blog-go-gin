package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog-go-gin/lib/db"
	"fmt"
	"strconv"
	"blog-go-gin/lib/paginator"
	"html/template"
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
	p := c.Param("page")
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}
	pageSize := 20
	offset := 0
	if page-1 > 0 {
		offset = pageSize * (page - 1)
	}

	Db := db.GetDb()

	countSql := "SELECT COUNT(*) as num FROM articles"
	totalRows := 0
	err = Db.Get(&totalRows, countSql)
	if err != nil {
		fmt.Printf("count article err:%#v\n", err)
		return
	}

	selectSql := fmt.Sprintf("SELECT * FROM articles LIMIT %d,%d", offset, pageSize)
	var arts []Article
	err = Db.Select(&arts, selectSql)
	if err != nil {
		fmt.Printf("select articles err:%#v\n", err)
		return
	}

	pageHtml := ""
	if totalRows > pageSize {
		pg := paginator.InitPages()
		pg.SetBaseUrl("/admin/article/index")
		pg.SetTotalRows(int32(totalRows))
		pg.SetCurrentPage(int32(page))
		pageHtml = pg.Create()
	}

	c.HTML(http.StatusOK, "admin/article_list.html", gin.H{
		"arts": arts,
		"pageHtml": template.HTML(pageHtml),
		"totalRows": totalRows,
	})
}

func ArticleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/article_add.html", gin.H{})
}

func ArticleEdit(c *gin.Context) {
	id := c.Param("id")
	sqlStr := "SELECT * FROM articles WHERE id=?"
	Db := db.GetDb()
	var art Article
	err := Db.Get(&art, sqlStr, id)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return
	}

	c.HTML(http.StatusOK, "admin/article_edit.html", gin.H{})
}

func ArticleDel(c *gin.Context) {
	id := c.PostForm("id")
	iId, _ := strconv.Atoi(id)
	if iId > 0 {
		sqlStr := "DELETE FROM articles WHERE id=?"
		Db := db.GetDb()
		_, err := Db.Exec(sqlStr, iId)
		if err != nil {
			fmt.Printf("del data err:%#v\n", err)
			c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "删除失败，请稍后重试！"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "删除成功！"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "参数错误！"})
}
