package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	r := gin.Default()

	r.POST("testUpload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//c.SaveUploadedFile(file, "./"+file.Filename) gin封装的 读取上传的文件流 目录新建一个本地文件 复制进去

		//原生的
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)

		//上传的文件传到前端洁面上
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
		c.File("./" + file.Filename)

		c.JSON(200, gin.H{
			"msg": file,
		})
	})

	r.Run()
}
