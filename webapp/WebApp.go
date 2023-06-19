package webapp

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

type WebApp struct {
}

func InitRouters(router *gin.Engine, app *WebApp) {
	router.SetFuncMap(template.FuncMap{
		"UnixToTime": app.UnixToTime,
		"UnixToDate": app.UnixToDate,
	})
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	G := router.Group("/", gin.BasicAuth(gin.Accounts{
		"lauson": "lauson",
	}))
	G.GET("/", app.Index)
	G.GET("/news", app.News)
	G.GET("/upload", app.Upload)
	G.POST("/doUpload", app.DoUpload)
	router.Run(":8080")
}
