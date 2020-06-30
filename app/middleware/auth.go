package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"encoding/json"
	"blog-go-gin/app/controllers/admin"
	//"fmt"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("admin_user")
		//fmt.Printf("sesion content: %#v\n", user)
		if user == nil {
			c.Abort()
			c.Redirect(http.StatusFound, "/admin/login")
		}
		userStr := user.(string)
		var users *admin.User
		err := json.Unmarshal([]byte(userStr), &users)
		if err == nil && users.Id > 0 {
			c.Next()
		} else {
			c.Abort()
			c.Redirect(http.StatusFound, "/admin/login")
		}
	}
}
