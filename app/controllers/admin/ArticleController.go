package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
	"blog-go-gin/lib/paginator"
	"html/template"
	"blog-go-gin/app/model/article"
	"blog-go-gin/app/model/category"
)

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

	totalRows := article.GetCount()

	arts := article.GetList(offset, pageSize)

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
	cate := category.GetList(true)
	c.HTML(http.StatusOK, "admin/article_add.html", gin.H{
		"cate": cate,
	})
}

func ArticleEdit(c *gin.Context) {
	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("convert string id to int err: %#v\n", err)
		return
	}

	art := article.GetOne(cId)

	_ = art

	c.HTML(http.StatusOK, "admin/article_edit.html", gin.H{})
}

func ArticleDel(c *gin.Context) {
	id := c.PostForm("id")
	iId, _ := strconv.Atoi(id)
	if iId > 0 {
		err := article.DeleteData(iId)
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
