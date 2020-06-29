package routes

import (
	"github.com/gin-gonic/gin"
	"blog-go-gin/app/controllers/admin"
	"blog-go-gin/app/controllers/web"
	"blog-go-gin/app/controllers"
	"blog-go-gin/app/middleware"
)

func WebRoutes(r *gin.Engine) {

	// 配置模板目录
	r.LoadHTMLGlob("app/views/**/*")

	// 配置静态资源目录
	r.Static("/public", "./public")

	// 验证码
	r.GET("/captcha", func(c *gin.Context) {
		controllers.Captcha(c,4)
	})

	// 后台
	r.GET("/admin/login", admin.Login)
	r.POST("/admin/login", admin.DoLogin)

	adminRoute := r.Group("/admin", middleware.CheckAuth())

	// 首页
	adminRoute.GET("/", admin.Index)
	adminRoute.GET("/index", admin.Index)

	// 退出
	adminRoute.GET("/logout", admin.Logout)

	// 栏目管理
	adminRoute.GET("/category/index", admin.CateIndex)
	adminRoute.GET("/category/add", admin.CateAdd)
	adminRoute.GET("/category/edit/:id", admin.CateEdit)
	adminRoute.GET("/category/del/:id", admin.CateDel)

	// 文章管理
	adminRoute.GET("/article/index", admin.ArticleIndex)
	adminRoute.GET("/article/add", admin.ArticleAdd)
	adminRoute.GET("/article/edit/:id", admin.ArticleEdit)
	adminRoute.GET("/article/del/:id", admin.ArticleDel)

	// 前台
	r.GET("/", web.Index)
	r.GET("/category/:id", web.Category)
	r.GET("/article/:id", web.Article)
	r.GET("/about", web.About)
	r.GET("/share", web.Share)
	r.GET("/share/:id", web.ShareInfo)

}
