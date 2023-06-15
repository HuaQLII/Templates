package webapp

import (
	"github.com/gin-gonic/gin"
)

type WebApp struct{
	JSCRC, CSSCRC string
	WebPort       string
	Nats_WSPort   string
}
type Article struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (this *WebApp) InitRouters(r *gin.Engine) {
	G := r.Group("/", gin.BasicAuth(gin.Accounts{
		"lauson": "lauson",
	}))
	_ = G
	G.GET("/", this.Index)
	r.Run(":8080")
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

}
