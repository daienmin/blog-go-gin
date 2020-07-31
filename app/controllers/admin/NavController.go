package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
	"blog-go-gin/app/model/nav"
)

func NavIndex(c *gin.Context)  {
	list := nav.GetList()
	c.HTML(http.StatusOK, "admin/nav_list.html", gin.H{
		"list": list,
	})
}

func NavAdd(c *gin.Context)  {
	c.HTML(http.StatusOK, "admin/nav_add.html", gin.H{})
}

func NavCreate(c *gin.Context)  {
	name := c.PostForm("name")
	url := c.PostForm("url")
	sort := c.PostForm("sort")
	target := c.PostForm("target")

	err := nav.InsertData(name, url, sort, target)
	if err != nil {
		fmt.Printf("insert data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "添加导航失败，请稍后重试！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "添加导航成功！"})
}

func NavEdit(c *gin.Context)  {
	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("convert string id to int err:%#v\n", err)
		return
	}
	data := nav.GetOne(cId)
	c.HTML(http.StatusOK, "admin/nav_edit.html", gin.H{
		"data": data,
	})
}

func NavUpdate(c *gin.Context)  {
	id := c.Param("id")
	name := c.PostForm("name")
	url := c.PostForm("url")
	sort := c.PostForm("sort")
	target := c.PostForm("target")

	err := nav.UpdateData(name, url, sort, id, target)

	if err != nil {
		fmt.Printf("update data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "修改导航失败，请稍后重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "修改导航成功！"})
}

func NavDel(c *gin.Context)  {
	id := c.Param("id")
	iId, _ := strconv.Atoi(id)
	if iId > 0 {
		err := nav.DeleteData(iId)
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