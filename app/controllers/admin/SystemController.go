package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog-go-gin/app/model/system"
)

func SystemSet(c *gin.Context) {
	conf := system.GetConfig()
	c.HTML(http.StatusOK, "admin/system_index.html", gin.H{
		"conf": conf,
	})
}

func SystemSave(c *gin.Context) {
	data := map[string]string{
		"c_title":       "",
		"c_keywords":    "",
		"c_description": "",
		"c_copy_icp":    "",
		"c_stat_code":   "",
	}
	data["c_title"] = c.PostForm("c_title")
	data["c_keywords"] = c.PostForm("c_keywords")
	data["c_description"] = c.PostForm("c_description")
	data["c_copy_icp"] = c.PostForm("c_copy_icp")
	data["c_stat_code"] = c.PostForm("c_stat_code")
	system.SaveData(data)
	c.Redirect(http.StatusFound, "/admin/system")
}
