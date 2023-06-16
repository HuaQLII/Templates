package webapp

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

type WebApp struct {
}

func InitRouters(router *gin.Engine, app *WebApp) {
	//自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"UnixToTime": app.UnixToTime,
		"UnixToDate": app.UnixToDate,

	})
	

	
	G := router.Group("/", gin.BasicAuth(gin.Accounts{
		"lauson": "lauson",
	}))
	G.GET("/", app.Index)
	G.GET("/news", app.News)
	router.Run(":8080")
}
