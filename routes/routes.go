package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathas/home-automation-helpers/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/triggers", controllers.CallVoiceMonkeyPost)
	r.GET("/triggers", controllers.CallVoiceMonkeyGet)
	r.NoRoute((controllers.EndpointNotFound))

	return r
}