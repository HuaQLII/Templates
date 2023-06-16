package controller

import (
	"Gin/model"
	"time"

	"github.com/gin-gonic/gin"
)

func GetNewsData() gin.H {
	articles := []interface{}{
		&model.Article{
			Name:    "新闻1",
			Address: "http://www.baidu.com",
		},
		&model.Article{
			Name:    "新闻2",
			Address: "http://www.baidu.com",
		},
	}

	data := gin.H{
		"title":   "新闻",
		"score":   100,
		"content": articles,
		"time":    time.Now().Unix(),
	}

	return data
}
