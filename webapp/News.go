package webapp

import (
	"Gin/controller"

	"github.com/gin-gonic/gin"
)

func (app *WebApp) News(c *gin.Context) {
	data := controller.GetNewsData()
	c.HTML(200, "default/news.html", data)
}

// r.GET("/news", func(c *gin.Context) {
	// 	news := &Article{
	// 		Name:    "新闻",
	// 		Address: "http://www.baidu.com",
	// 	}

	// 	c.HTML(200, "default/news.html", gin.H{
	// 		"title": "新闻",
	// 		"news":  news,
	// 	})
	// })

	// //后台

	// r.GET("/admin", func(c *gin.Context) {
	// 	c.HTML(200, "admin/index.html", gin.H{
	// 		"title": "后台首页",
	// 	})
	// })
	// r.GET("/admin/news", func(c *gin.Context) {
	// 	c.HTML(200, "admin/news.html", gin.H{
	// 		"title": "后台新闻",
	// 	})
	// })
