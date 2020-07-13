package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog-go-gin/app/model/article"
	"blog-go-gin/app/model/nav"
)

func Index(c *gin.Context) {
	list := article.GetIndexList(15)
	tagStat := article.GetTagsStat()
	hotViews := article.GetArtByParam(0)
	topList := article.GetArtByParam(1)
	recommendList := article.GetArtByParam(2)
	navs := nav.GetNavs()
	c.HTML(http.StatusOK, "home1/index.html", gin.H{
		"list":          list,
		"tagStat":       tagStat,
		"hotViews":      hotViews,
		"topList":       topList,
		"recommendList": recommendList,
		"navs":          navs,
	})
}

func Category(c *gin.Context) {
	c.HTML(http.StatusOK, "home1/list.html", gin.H{})
}

func Article(c *gin.Context) {
	id := c.Param("id")
	info := article.GetArtInfo(id)
	tagStat := article.GetTagsStat()
	hotViews := article.GetArtByParam(0)
	topList := article.GetArtByParam(1)
	recommendList := article.GetArtByParam(2)
	navs := nav.GetNavs()
	c.HTML(http.StatusOK, "home1/info.html", gin.H{
		"info":          info,
		"tagStat":       tagStat,
		"hotViews":      hotViews,
		"topList":       topList,
		"recommendList": recommendList,
		"navs":          navs,
	})
}

func TagList(c *gin.Context) {
	c.HTML(http.StatusOK, "home1/list.html", gin.H{})
}
