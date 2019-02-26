package routes

import (
	"github.com/blitzkriegcoding/todo-gin-mgo/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRouter(router *gin.RouterGroup) {
	router.POST("/", controllers.CreateTodo)
	router.GET("/", controllers.FetchAllTodo)
	router.GET("/:id", controllers.FetchSingleTodo)
	router.PUT("/:id", controllers.UpdateTodo)
	router.DELETE("/:id", controllers.DeleteTodo)
}
