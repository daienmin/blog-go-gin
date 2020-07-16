package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
	"blog-go-gin/app/model/adminuser"
	"blog-go-gin/lib/helper"
	"blog-go-gin/config"
)

func AdminUserIndex(c *gin.Context)  {
	list := adminuser.GetList()
	c.HTML(http.StatusOK, "admin/admin_user_list.html", gin.H{
		"list": list,
	})
}

func AdminUserAdd(c *gin.Context)  {
	c.HTML(http.StatusOK, "admin/admin_user_add.html", gin.H{})
}

func AdminUserCreate(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := adminuser.InsertData(username, helper.AesEncrypt(password, config.GetCfg().AppKey))
	if err != nil {
		fmt.Printf("insert data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "添加后台用户失败，请稍后重试！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "添加后台用户成功！"})
}

func AdminUserEdit(c *gin.Context)  {
	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("convert string id to int err:%#v\n", err)
		return
	}
	data := adminuser.GetOne(cId)
	c.HTML(http.StatusOK, "admin/admin_user_edit.html", gin.H{
		"data": data,
	})
}

func AdminUserUpdate(c *gin.Context)  {
	id := c.Param("id")
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := adminuser.UpdateData(username, helper.AesEncrypt(password, config.GetCfg().AppKey), id)

	if err != nil {
		fmt.Printf("update data err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": 1, "msg": "修改后台用户失败，请稍后重试！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "修改后台用户成功！"})
}

func AdminUserDel(c *gin.Context)  {
	id := c.Param("id")
	iId, _ := strconv.Atoi(id)
	if iId > 0 {
		err := adminuser.DeleteData(iId)
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