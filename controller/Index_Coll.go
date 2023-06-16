package controller

import (
	"Gin/model"

	"github.com/gin-gonic/gin"
)

func GetIndexData() gin.H {
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
		"title":   "首页",
		"score":   100,
		"content": articles,
	}

	return data
}
