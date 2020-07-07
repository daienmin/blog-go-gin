package admin

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"blog-go-gin/app/controllers"
	"blog-go-gin/lib/db"
	"blog-go-gin/lib/helper"
	"blog-go-gin/config"
	"github.com/gin-contrib/sessions"
	"encoding/json"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.html", gin.H{})
}

type LoginForm struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

type User struct {
	Id       uint32 `db:"id" json:"id"`
	UserName string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

func DoLogin(c *gin.Context) {

	var form LoginForm
	if err := c.Bind(&form); err != nil {
		fmt.Printf("param bind err: %#v\n", err.Error());
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "参数错误！"})
		return
	}

	// 验证验证码
	if flag := controllers.CaptchaVerify(c, form.Code); !flag {
		fmt.Printf("captcha code error\n");
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "验证码错误！"})
		return
	}

	// 获取用户信息
	Db := db.GetDb()
	var user User
	sqlStr := "select id,username,password from admin_users where username=?"
	err := Db.Get(&user, sqlStr, form.UserName)
	if err != nil {
		fmt.Printf("get user info err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "用户或密码错误！"})
		return
	}

	// 验证密码是否正确
	if helper.AesDecrypt(user.Password, config.GetCfg().AppKey) != form.Password {
		fmt.Printf("user input password err", )
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "用户或密码错误！"})
		return
	}

	// 保存用户信息到session
	session := sessions.Default(c)
	byteUser, _ := json.Marshal(user)
	session.Set("admin_user", string(byteUser))
	_ = session.Save()

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登录成功！"})
}

func Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Delete("admin_user")
	session.Save()
	c.Redirect(http.StatusFound, "/admin/login")
}