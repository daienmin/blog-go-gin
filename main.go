package main

import (
	"blog-go-gin/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"blog-go-gin/config"
	_ "github.com/gorilla/sessions"
)

func main() {

	// 创建gin框架实例
	r := gin.Default()

	// 加载路由
	routes.WebRoutes(r)

	// 加载配置文件
	err := config.LoadCfg("config/config.yaml")
	if err != nil {
		fmt.Printf("err to load config: %v\n", err)
		return
	}
	
	// 加载监听端口，启动server
	err = r.Run(config.GetCfg().Port)
	if err != nil {
		fmt.Printf("start server faild: %v\n", err)
		return
	}
}
