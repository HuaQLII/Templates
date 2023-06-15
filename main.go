package main

import (
	"Gin/db"
	"Gin/webapp"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.InitDB()
	app := &webapp.WebApp{} // 创建 WebApp 的实例
	app.InitRouters(r) // 调用 InitRouters 方法
}
