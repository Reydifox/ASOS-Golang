package api

import (
	"ASOS/models"
	"ASOS/pkg/app"
	"ASOS/pkg/flags"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	appGin := app.Gin{C: c}

	appGin.Response(http.StatusOK, flags.SUCCESS, "Hello ASOS!")
}

func GetExample(c *gin.Context) {
	appGin := app.Gin{C: c}
	name := c.Param("name")

	example := models.GetExample(name)
	if example == nil {
		appGin.Response(http.StatusOK, flags.ERROR_NOT_EXIST_EXAMPLE, nil)
		return
	}

	appGin.Response(http.StatusOK, flags.SUCCESS, example)
}

func GetMessages(c *gin.Context) {
	appGin := app.Gin{C: c}
	name := c.Param("name")

	example := models.GetExample(name)
	if example == nil {
		appGin.Response(http.StatusOK, flags.ERROR_NOT_EXIST_EXAMPLE, nil)
		return
	}

	appGin.Response(http.StatusOK, flags.SUCCESS, example.Messages)
}

func GetMessage(c *gin.Context) {
	appGin := app.Gin{C: c}
	name := c.Param("name")
	message := c.Param("message")

	example := models.GetExample(name)
	if example == nil {
		appGin.Response(http.StatusOK, flags.ERROR_NOT_EXIST_EXAMPLE, nil)
		return
	}

	msgs := models.GetMessage(example, message)
	if msgs == nil {
		appGin.Response(http.StatusOK, flags.ERROR_NOT_EXIST_MESSAGE, nil)
		return
	}

	appGin.Response(http.StatusOK, flags.SUCCESS, msgs)
}
