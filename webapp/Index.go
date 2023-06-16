package webapp

import (
	"Gin/controller"

	"github.com/gin-gonic/gin"
)

func (app *WebApp) Index(c *gin.Context) {
	data := controller.GetIndexData()
	c.HTML(200, "default/index.html", data)
}
