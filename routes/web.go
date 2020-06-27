package routes

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"blog-go-gin/app/controllers/admin"
	_ "github.com/gorilla/sessions"
)


func WebRoutes(r *gin.Engine) {

	r.LoadHTMLFiles()
	// 后台
	adminRoute := r.Group("/admin")

	adminRoute.GET("/", admin.Index)
	adminRoute.GET("/index", admin.Index)


}
