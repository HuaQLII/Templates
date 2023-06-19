package main

import (
	"Gin/db"
	"Gin/webapp"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	//自定义模板函数
	app := &webapp.WebApp{}
	webapp.InitRouters(router, app)
}
