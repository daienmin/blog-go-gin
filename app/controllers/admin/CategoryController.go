package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
	"blog-go-gin/app/model/category"
)

func CateIndex(c *gin.Context) {
	categories := category.GetList(false)
	c.HTML(http.StatusOK, "admin/cate_list.html", gin.H{"categories": categories})
}

func CateAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/cate_add.html", gin.H{})
}

func CateCreate(c *gin.Context) {
	pid := c.PostForm("pid")
	name := c.PostForm("name")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort := c.PostForm("sort")
	status := c.PostForm("status")

	err := category.InsertData(pid, name, keywords, description, sort, status)
	if err != nil {
		fmt.Printf("insert data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "添加栏目失败，请稍后重试！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "添加栏目成功！"})
}

func CateEdit(c *gin.Context) {
	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("convert string id to int err:%#v\n", err)
		return
	}
	cate := category.GetOne(cId)

	c.HTML(http.StatusOK, "admin/cate_edit.html", gin.H{"category": cate})
}

func CateUpdate(c *gin.Context) {
	id := c.Param("id")
	pid := c.PostForm("pid")
	name := c.PostForm("name")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort := c.PostForm("sort")
	status := c.PostForm("status")

	err := category.UpdateData(pid, name, keywords, description, sort, status, id)

	if err != nil {
		fmt.Printf("update data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "修改栏目失败，请稍后重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "修改栏目成功！"})
}

func CateDel(c *gin.Context) {
	id := c.PostForm("id")
	iId, _ := strconv.Atoi(id)
	if iId > 0 {
		err := category.DeleteData(iId)
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
