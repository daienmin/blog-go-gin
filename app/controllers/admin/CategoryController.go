package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"blog-go-gin/app/helper"
)

type Category struct {
	Id          uint32 `db:"id" json:"id"`
	Pid         int32  `db:"pid" json:"pid"`
	Name        string `db:"name" json:"name"`
	KeyWords    string `db:"keywords" json:"keywords"`
	Description string `db:"description" json:"description"`
	Sort        uint8  `db:"sort" json:"sort"`
	Status      uint8  `db:"status" json:"status"`
	CreatedAt   string  `db:"created_at" json:"created_at"`
	UpdatedAt   string  `db:"updated_at" json:"updated_at"`
}

func CateIndex(c *gin.Context) {
	sqlStr := "SELECT * FROM categories"
	db := helper.GetDb()
	var categories []Category
	err := db.Select(&categories, sqlStr)
	if err != nil {
		fmt.Printf("select err:%#v\n", err)
		return
	}
	fmt.Printf("categories data: %#v\n", categories)

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
	createdAt := helper.GetDateTime()
	updatedAt := createdAt
	db := helper.GetDb()

	sqlStr := "INSERT INTO categories(pid,name,keywords,description,sort,status,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?)"

	_, err := db.Exec(sqlStr, pid, name, keywords, description, sort, status, createdAt, updatedAt)
	if err != nil {
		fmt.Printf("insert data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "添加栏目失败，请稍后重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": ""})
}

func CateEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/cate_edit.html", gin.H{})
}

func CateDel(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/cate_add.html", gin.H{})
}
