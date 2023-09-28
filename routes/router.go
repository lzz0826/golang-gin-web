package routes

import (
	"github.com/gin-gonic/gin"
	"web/controllers"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", controllers.IndexHome)
	router.GET("/index", controllers.IndexHome)

	router.GET("/users/:id", controllers.UserGet)
	router.GET("/users", controllers.UserGetList)
	router.POST("/users", controllers.UserPost)
	router.PUT("/users/:id", controllers.UserPut)
	router.PATCH("/users/:id", controllers.UserPut)
	router.DELETE("/users/:id", controllers.UserDelete)
}
