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
	"blog-go-gin/app/model/tag"
	"encoding/json"
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
		pg.SetListRows(int32(pageSize))
		pageHtml = pg.Create()
	}

	c.HTML(http.StatusOK, "admin/article_list.html", gin.H{
		"arts":      arts,
		"pageHtml":  template.HTML(pageHtml),
		"totalRows": totalRows,
	})
}

func ArticleAdd(c *gin.Context) {
	cate := category.GetList(true)
	tags := tag.GetList()
	c.HTML(http.StatusOK, "admin/article_add.html", gin.H{
		"cate": cate,
		"tags": tags,
	})
}

func ArticleCreate(c *gin.Context) {
	title := c.PostForm("title")
	categoryId := c.PostForm("category_id")
	keywords := c.PostForm("keywords")
	cover := c.PostForm("cover")
	description := c.PostForm("description")
	markdown := c.PostForm("markdown")
	html := c.PostForm("html")
	status := c.PostForm("status")
	isRecommend := c.PostForm("is_recommend")
	isTop := c.PostForm("is_top")
	author := c.PostForm("author")
	tagIds := c.PostFormArray("tag_ids[]")

	cateId, _ := strconv.Atoi(categoryId)
	stat, _ := strconv.Atoi(status)
	isRec, _ := strconv.Atoi(isRecommend)
	isT, _ := strconv.Atoi(isTop)
	artData := &article.Article{
		Title:       title,
		CategoryId:  uint8(cateId),
		KeyWords:    keywords,
		Cover:       cover,
		Description: description,
		MarkDown:    markdown,
		Html:        html,
		Author:      author,
		Status:      uint8(stat),
		IsRecommend: uint8(isRec),
		IsTop:       uint8(isT),
	}
	err := article.InsertData(artData, tagIds)
	if err != nil {
		fmt.Printf("insert article data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"error": 1,
			"msg": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": 0,
		"msg": "添加成功",
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
	cate := category.GetList(true)
	tags := tag.GetList()
	artTags := article.GetTags(art.Id)
	at, err := json.Marshal(artTags)
	if err != nil {
		fmt.Printf("marshal artTags err:%#v\n", err)
	}
	c.HTML(http.StatusOK, "admin/article_edit.html", gin.H{
		"art": art,
		"cate": cate,
		"tags": tags,
		"artTags":string(at),
	})
}

func ArticleUpdate(c *gin.Context)  {
	id := c.Param("id")
	title := c.PostForm("title")
	categoryId := c.PostForm("category_id")
	keywords := c.PostForm("keywords")
	cover := c.PostForm("cover")
	description := c.PostForm("description")
	markdown := c.PostForm("markdown")
	html := c.PostForm("html")
	status := c.PostForm("status")
	isRecommend := c.PostForm("is_recommend")
	isTop := c.PostForm("is_top")
	author := c.PostForm("author")
	tagIds := c.PostFormArray("tag_ids[]")

	cateId, _ := strconv.Atoi(categoryId)
	stat, _ := strconv.Atoi(status)
	isRec, _ := strconv.Atoi(isRecommend)
	isT, _ := strconv.Atoi(isTop)
	artData := &article.Article{
		Title:       title,
		CategoryId:  uint8(cateId),
		KeyWords:    keywords,
		Cover:       cover,
		Description: description,
		MarkDown:    markdown,
		Html:        html,
		Author:      author,
		Status:      uint8(stat),
		IsRecommend: uint8(isRec),
		IsTop:       uint8(isT),
	}
	err := article.UpdateData(artData, tagIds, id)
	if err != nil {
		fmt.Printf("update article data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"error": 1,
			"msg": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": 0,
		"msg": "修改成功",
	})
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
