package webapp

import "github.com/gin-gonic/gin"

func (this *WebApp) Index(c *gin.Context) {
	c.HTML(200, "default/index.html", gin.H{
		"title": "首页",
		"score": 100,
		"content": []interface{}{
			&Article{
				Name:    "新闻1",
				Address: "http://www.baidu.com",
			},
			&Article{
				Name:    "新闻2",
				Address: "http://www.baidu.com",
			},
		},
	})
}
