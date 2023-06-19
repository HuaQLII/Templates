package webapp

import (
	"github.com/gin-gonic/gin"
)

func (app *WebApp) Upload(c *gin.Context) {
	c.HTML(200, "upload.html", gin.H{
		"title": "上传",
		"code":  200,
	})
}

func (app *WebApp) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	file, e := c.FormFile("file")
	if e == nil {
		c.SaveUploadedFile(file, "upload/"+file.Filename)
	}
	// c.String(200, "上传成功,用户名为:%s,文件名为:%s", username, file.Filename)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"username": username,
	})
}