package routers

import (
	"ASOS/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	restapi := r.Group("/api")
	restapi.GET("/example", api.GetInfo)
	restapi.GET("/example/:name", api.GetExample)
	restapi.GET("/example/:name/messages/", api.GetMessages)
	restapi.GET("/example/:name/messages/:message", api.GetMessage)

	return r
}
