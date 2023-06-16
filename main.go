package main

import (
	"Gin/db"
	"Gin/webapp"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static")
	app := &webapp.WebApp{}
	webapp.InitRouters(router, app)
}
