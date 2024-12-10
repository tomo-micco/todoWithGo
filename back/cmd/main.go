package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomo-micco/TodoWithGo/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/todo", handlers.GetAllTodo)
	r.GET("/todo/:id", handlers.FindTodoById)
	r.POST("/todo", handlers.Create)
	r.PUT("/todo", handlers.Update)
	r.DELETE("/todo/:id", handlers.Delete)
	r.Run()
}
