package main

import (
	"blog-go-gin/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"blog-go-gin/config"
	"blog-go-gin/lib/db"
	"blog-go-gin/lib/sess"
	"blog-go-gin/lib/redis_pool"
)

func main() {

	// 加载配置文件
	err := config.LoadCfg("config/config.yaml")
	if err != nil {
		fmt.Printf("err to load config: %v\n", err)
		return
	}

	// 初始化数据库连接
	err = db.InitDb()
	if err != nil {
		fmt.Printf("err to init db: %#v\n", err)
		return
	}

	cfg := config.GetCfg()

	// 初始化redis链接池
	err = redis_pool.InitPool(cfg.RedisHost + ":" + cfg.RedisPort, cfg.RedisPassword)
	if err != nil {
		fmt.Printf("err to init redis pool:%#v\n", err)
		return
	}

	// 创建gin框架实例
	r := gin.Default()

	// 初始化session
	r.Use(sess.Session(cfg.AppKey))

	// 加载路由
	routes.WebRoutes(r)

	// 加载监听端口，启动server
	err = r.Run(config.GetCfg().Port)
	if err != nil {
		fmt.Printf("start server faild: %v\n", err)
		return
	}
}
