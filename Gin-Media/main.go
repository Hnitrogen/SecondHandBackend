package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 支持的图片类型
var allowedTypes = map[string]bool{
	"avatar": true,
	"stuff":  true,
}

func main() {
	createStorageDirs()

	r := gin.Default()

	// 使用 cors 中间件
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// 或者指定域名
	// config.AllowOrigins = []string{"http://localhost:8080"}
	r.Use(cors.New(config))

	r.POST("media/upload", handleUpload)
	r.GET("media/image/:type/:filename", handleGetImage)

	r.Run(":8080")
}

func createStorageDirs() {
	os.MkdirAll("storage/avatar", 0755)
	os.MkdirAll("storage/stuff", 0755)
}

func handleUpload(c *gin.Context) {
	// 获取图片类型
	imageType := c.PostForm("type")
	if !allowedTypes[imageType] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "不支持的图片类型",
		})
		return
	}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "文件上传失败",
		})
		return
	}

	// 生成文件哈希名
	hashedName, err := saveFile(file, imageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type":     imageType,
		"filename": hashedName,
		"preview":  "http://localhost/media/image/" + imageType + "/" + hashedName,
	})
}

func saveFile(file *multipart.FileHeader, imageType string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 计算文件MD5哈希
	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		return "", err
	}
	src.Seek(0, 0)

	// 使用原始文件扩展名
	ext := path.Ext(file.Filename)
	hashedName := fmt.Sprintf("%x%s", hash.Sum(nil), ext)

	// 构建存储路径
	dst := fmt.Sprintf("./storage/%s/%s", imageType, hashedName)

	// 创建目标文件
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 保存文件
	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	return hashedName, nil
}

func handleGetImage(c *gin.Context) {
	imageType := c.Param("type")
	filename := c.Param("filename")

	if !allowedTypes[imageType] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "不支持的图片类型",
		})
		return
	}

	filepath := fmt.Sprintf("./storage/%s/%s", imageType, filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "图片不存在",
		})
		return
	}

	c.File(filepath)
}
