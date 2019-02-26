package main

import (
	"github.com/blitzkriegcoding/todo-gin-mgo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	routes.TodoRouter(v1.Group("/todos"))
	r.Run(":4001")
}
