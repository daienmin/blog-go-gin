package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
	"blog-go-gin/app/model/tag"
)

func TagIndex(c *gin.Context) {
	tags := tag.GetList()
	fmt.Printf("tags data: %#v\n", tags)
	c.HTML(http.StatusOK, "admin/tag_list.html", gin.H{"tags": tags})
}

func TagAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/tag_add.html", gin.H{})
}

func TagCreate(c *gin.Context) {
	name := c.PostForm("name")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	err := tag.InsertData(name, keywords, description)
	if err != nil {
		fmt.Printf("insert data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "添加标签失败，请稍后重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "添加标签成功！"})
}

func TagEdit(c *gin.Context) {
	id := c.Param("id")
	cId,_ := strconv.Atoi(id)
	tagData := tag.GetOne(cId)
	c.HTML(http.StatusOK, "admin/tag_edit.html", gin.H{"tag": tagData})
}

func TagUpdate(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	err := tag.UpdateData(name, keywords, description, id)
	if err != nil {
		fmt.Printf("update data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "修改标签失败，请稍后重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "修改标签成功！"})
}

func TagDel(c *gin.Context) {
	id := c.PostForm("id")
	iId, _ := strconv.Atoi(id)
	if iId > 0 {
		err := tag.DeleteData(iId)
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
