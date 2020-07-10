package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog-go-gin/app/model/link"
	"fmt"
	"strconv"
)

func LinkIndex(c *gin.Context)  {
	list := link.GetList()
	c.HTML(http.StatusOK, "admin/link_list.html", gin.H{
		"list": list,
	})
}

func LinkAdd(c *gin.Context)  {
	c.HTML(http.StatusOK, "admin/link_add.html", gin.H{})
}

func LinkCreate(c *gin.Context)  {
	name := c.PostForm("name")
	url := c.PostForm("url")
	sort := c.PostForm("sort")

	err := link.InsertData(name, url, sort)
	if err != nil {
		fmt.Printf("insert data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "添加友情链接失败，请稍后重试！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "添加友情链接成功！"})
}

func LinkEdit(c *gin.Context)  {
	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("convert string id to int err:%#v\n", err)
		return
	}
	data := link.GetOne(cId)
	c.HTML(http.StatusOK, "admin/link_edit.html", gin.H{
		"data": data,
	})
}

func LinkUpdate(c *gin.Context)  {
	id := c.Param("id")
	name := c.PostForm("name")
	url := c.PostForm("url")
	sort := c.PostForm("sort")

	err := link.UpdateData(name, url, sort, id)

	if err != nil {
		fmt.Printf("update data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "修改友情链接失败，请稍后重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "修改友情链接成功！"})
}

func LinkDel(c *gin.Context)  {
	id := c.Param("id")
	iId, _ := strconv.Atoi(id)
	if iId > 0 {
		err := link.DeleteData(iId)
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