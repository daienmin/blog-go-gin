package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"blog-go-gin/lib/helper"
	"os"
	"strings"
	"math/rand"
	"time"
	"strconv"
)

// 图片上传
func ArticleUploadImg(c *gin.Context) {
	_, h, err := c.Request.FormFile("editormd-image-file")
	if err != nil {
		fmt.Printf("upload file error: %#v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"success": 0,
			"message": "图片上传失败",
		})
		return
	}
	// 上传路径
	filePath := "./public/upload/articles/" + helper.GetDate() + "/"
	// 创建目录
	os.Mkdir(filePath, os.ModeDir)

	//headers.Size 获取文件大小
	/*if h.Size > 1024*1024*2 {
		fmt.Println("文件太大了")
		return
	}*/
	//headers.Header.Get("Content-Type")获取上传文件的类型
	/*if h.Header.Get("Content-Type") != "image/png" {
		fmt.Println("只允许上传png图片")
		return
	}*/

	// 获取文件扩展名
	fArr := strings.Split(h.Filename, ".")
	ext := fArr[len(fArr) - 1]
	// 生成随机文件名
	rand.Seed(time.Now().UnixNano())
	name := rand.Intn(9999)
	filename := helper.GetDateTimeCom() + strconv.Itoa(name)
	url := filePath + filename + "." + ext
	// 保存文件
	err = c.SaveUploadedFile(h, url)
	if err != nil {
		fmt.Printf("save file err:%#v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"success": 0,
			"message": "保存文件失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": 1,
		"message": "上传成功",
		"url": strings.Trim(url, "."),
	})
}