package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog-go-gin/app/model/article"
	"blog-go-gin/app/model/nav"
	"html/template"
	"strconv"
	"blog-go-gin/lib/paginator"
	"blog-go-gin/app/model/category"
	"blog-go-gin/app/model/tag"
)

var (
	tagStat       []article.TagStat
	hotViews      []article.Article
	topList       []article.Article
	recommendList []article.Article
	navs          []nav.Nav
	allCate       []category.Category
)

func common() {
	tagStat = article.GetTagsStat()
	hotViews = article.GetArtByParam(0)
	topList = article.GetArtByParam(1)
	recommendList = article.GetArtByParam(2)
	navs = nav.GetNavs()
	allCate = category.GetList(true)
}

// 首页
func Index(c *gin.Context) {
	list := article.GetIndexList(15)

	common()
	c.HTML(http.StatusOK, "home1/index.html", gin.H{
		"list":          list,
		"tagStat":       tagStat,
		"hotViews":      hotViews,
		"topList":       topList,
		"recommendList": recommendList,
		"navs":          navs,
		"allCate":       allCate,
	})
}

// 分类列表页
func Category(c *gin.Context) {
	id := c.Param("id")
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
	totalRows := article.GetCateArtCount(id)
	list := article.GetCateArtList(id, offset, pageSize)
	cateId, _ := strconv.Atoi(id)
	cate := category.GetOne(cateId)

	pageHtml := ""
	if totalRows > pageSize {
		pg := paginator.InitPages()
		pg.SetBaseUrl("/article/" + id)
		pg.SetTotalRows(int32(totalRows))
		pg.SetCurrentPage(int32(page))
		pg.SetListRows(int32(pageSize))
		pageHtml = pg.Create()
	}

	common()
	c.HTML(http.StatusOK, "home1/list.html", gin.H{
		"list":          list,
		"pageHtml":      template.HTML(pageHtml),
		"cate":          cate,
		"tagStat":       tagStat,
		"hotViews":      hotViews,
		"topList":       topList,
		"recommendList": recommendList,
		"navs":          navs,
		"allCate":       allCate,
	})
}

// 文章详情页
func Article(c *gin.Context) {
	id := c.Param("id")
	art := article.GetArtInfo(id)
	artTag := article.GetTags(art.Id)
	optArt := article.GetOptArt(art.Id)

	common()
	c.HTML(http.StatusOK, "home1/info.html", gin.H{
		"art":           art,
		"artContent":    template.HTML(art.Html),
		"artTag":        artTag,
		"optArt":        optArt,
		"tagStat":       tagStat,
		"hotViews":      hotViews,
		"topList":       topList,
		"recommendList": recommendList,
		"navs":          navs,
		"allCate":       allCate,
	})
}

// 标签列表页
func TagList(c *gin.Context) {
	id := c.Param("id")
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
	totalRows := article.GetTagArtCount(id)
	list := article.GetTagArtList(id, offset, pageSize)
	tagId, _ := strconv.Atoi(id)
	tagInfo := tag.GetOne(tagId)

	pageHtml := ""
	if totalRows > pageSize {
		pg := paginator.InitPages()
		pg.SetBaseUrl("/article/" + id)
		pg.SetTotalRows(int32(totalRows))
		pg.SetCurrentPage(int32(page))
		pg.SetListRows(int32(pageSize))
		pageHtml = pg.Create()
	}

	common()
	c.HTML(http.StatusOK, "home1/tag_list.html", gin.H{
		"list":          list,
		"pageHtml":      template.HTML(pageHtml),
		"tagInfo":       tagInfo,
		"tagStat":       tagStat,
		"hotViews":      hotViews,
		"topList":       topList,
		"recommendList": recommendList,
		"navs":          navs,
		"allCate":       allCate,
	})
}
