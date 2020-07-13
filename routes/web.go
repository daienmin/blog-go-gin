package routes

import (
	"github.com/gin-gonic/gin"
	"blog-go-gin/app/controllers/admin"
	"blog-go-gin/app/controllers/web"
	"blog-go-gin/app/controllers"
	"blog-go-gin/app/middleware"
	"html/template"
	"blog-go-gin/app/model/common"
)

func WebRoutes(r *gin.Engine) {

	// 注册模板函数
	r.SetFuncMap(template.FuncMap{
		"getArtCate": common.GetArtCate, // 获取文章分类名
		"getArtTags": common.GetArtTags, // 获取文章标签名
		"mod":        common.Mod,        // 获取文章标签名
	})

	// 配置模板目录
	r.LoadHTMLGlob("app/views/**/*")

	// 配置静态资源目录
	r.Static("/public", "./public")

	// 验证码
	r.GET("/captcha", func(c *gin.Context) {
		controllers.Captcha(c, 4)
	})

	// 前台
	r.GET("/", web.Index)
	r.GET("/category/:id", web.Category)
	r.GET("/article/:id", web.Article)
	r.GET("/tag/:id", web.TagList)

	// 后台
	r.GET("/admin/login", admin.Login)
	r.POST("/admin/login", admin.DoLogin)

	adminRoute := r.Group("/admin", middleware.CheckAuth())

	// 首页
	adminRoute.GET("/", admin.Index)
	adminRoute.GET("/index", admin.Index)

	// 退出
	adminRoute.GET("/logout", admin.Logout)

	// 后台上传
	adminRoute.POST("/upload/article", admin.ArticleUploadImg)

	// 栏目管理
	adminRoute.GET("/category/index", admin.CateIndex)
	adminRoute.GET("/category/add", admin.CateAdd)
	adminRoute.POST("/category/add", admin.CateCreate)
	adminRoute.GET("/category/edit/:id", admin.CateEdit)
	adminRoute.POST("/category/edit/:id", admin.CateUpdate)
	adminRoute.POST("/category/del", admin.CateDel)

	// 文章标签管理
	adminRoute.GET("/tag/index", admin.TagIndex)
	adminRoute.GET("/tag/add", admin.TagAdd)
	adminRoute.POST("/tag/add", admin.TagCreate)
	adminRoute.GET("/tag/edit/:id", admin.TagEdit)
	adminRoute.POST("/tag/edit/:id", admin.TagUpdate)
	adminRoute.POST("/tag/del", admin.TagDel)

	// 文章管理
	adminRoute.GET("/article/index", admin.ArticleIndex)
	adminRoute.GET("/article/index/:page", admin.ArticleIndex)
	adminRoute.GET("/article/add", admin.ArticleAdd)
	adminRoute.POST("/article/add", admin.ArticleCreate)
	adminRoute.GET("/article/edit/:id", admin.ArticleEdit)
	adminRoute.POST("/article/edit/:id", admin.ArticleUpdate)
	adminRoute.GET("/article/del/:id", admin.ArticleDel)

	// 导航管理
	adminRoute.GET("/nav/index", admin.NavIndex)
	//adminRoute.GET("/nav/index/:page", admin.NavIndex)
	adminRoute.GET("/nav/add", admin.NavAdd)
	adminRoute.POST("/nav/add", admin.NavCreate)
	adminRoute.GET("/nav/edit/:id", admin.NavEdit)
	adminRoute.POST("/nav/edit/:id", admin.NavUpdate)
	adminRoute.GET("/nav/del/:id", admin.NavDel)

	// 友情链接管理
	adminRoute.GET("/link/index", admin.LinkIndex)
	//adminRoute.GET("/link/index/:page", admin.LinkIndex)
	adminRoute.GET("/link/add", admin.LinkAdd)
	adminRoute.POST("/link/add", admin.LinkCreate)
	adminRoute.GET("/link/edit/:id", admin.LinkEdit)
	adminRoute.POST("/link/edit/:id", admin.LinkUpdate)
	adminRoute.GET("/link/del/:id", admin.LinkDel)

	// 后台用户管理
	adminRoute.GET("/admin_user/index", admin.AdminUserIndex)
	//adminRoute.GET("/admin_user/index/:page", admin.AdminUserIndex)
	adminRoute.GET("/admin_user/add", admin.AdminUserAdd)
	adminRoute.POST("/admin_user/add", admin.AdminUserCreate)
	adminRoute.GET("/admin_user/edit/:id", admin.AdminUserEdit)
	adminRoute.POST("/admin_user/edit/:id", admin.AdminUserUpdate)
	adminRoute.GET("/admin_user/del/:id", admin.AdminUserDel)

}
