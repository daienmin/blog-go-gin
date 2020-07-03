package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"blog-go-gin/app/helper"
	"strconv"
)

type Tag struct {
	Id          uint32 `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	KeyWords    string `db:"keywords" json:"keywords"`
	Description string `db:"description" json:"description"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
}

func TagIndex(c *gin.Context) {
	sqlStr := "SELECT * FROM tags"
	db := helper.GetDb()
	var tags []Tag
	err := db.Select(&tags, sqlStr)
	if err != nil {
		fmt.Printf("select err:%#v\n", err)
		return
	}
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
	createdAt := helper.GetDateTime()
	updatedAt := createdAt
	db := helper.GetDb()

	sqlStr := "INSERT INTO tags(name,keywords,description,created_at,updated_at) VALUES(?,?,?,?,?)"

	_, err := db.Exec(sqlStr, name, keywords, description, createdAt, updatedAt)
	if err != nil {
		fmt.Printf("insert data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "添加标签失败，请稍后重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "添加标签成功！"})
}

func TagEdit(c *gin.Context) {
	id := c.Param("id")
	sqlStr := "SELECT * FROM tags WHERE id=?"
	db := helper.GetDb()
	var tag Tag
	err := db.Get(&tag, sqlStr, id)
	if err != nil {
		fmt.Printf("query err:%#v\n", err)
		return
	}
	c.HTML(http.StatusOK, "admin/tag_edit.html", gin.H{"tag": tag})
}

func TagUpdate(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	createdAt := helper.GetDateTime()
	updatedAt := createdAt
	db := helper.GetDb()

	sqlStr := "UPDATE tags SET name=?,keywords=?,description=?,updated_at=? WHERE id=?"

	_, err := db.Exec(sqlStr, name, keywords, description, updatedAt, id)
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
		sqlStr := "DELETE FROM tags WHERE id=?"
		db := helper.GetDb()
		_, err := db.Exec(sqlStr, iId)
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
