package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomo-micco/TodoWithGo/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/todo", handlers.GetAllTodo)
	r.GET("/todo/:id", handlers.FindTodoById)
	r.Run()
}
