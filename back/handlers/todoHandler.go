package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tomo-micco/TodoWithGo/databases/repositories"
	"github.com/tomo-micco/TodoWithGo/infrastructure"
	"github.com/tomo-micco/TodoWithGo/useCases"
)

/*
* 全件取得
 */
func GetAllTodo(c *gin.Context) {
	db, err := infrastructure.NewDB()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	defer db.Close()

	repository := repositories.NewTodoRepository(db)
	useCase := useCases.NewGetTodoUseCase(repository)
	todos, err := useCase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errorMessage": err,
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

/*
* IDに該当するTodo取得
 */
func FindTodoById(c *gin.Context) {
	db, err := infrastructure.NewDB()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	defer db.Close()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is not number",
		})
	}
	repository := repositories.NewTodoRepository(db)
	useCases := useCases.NewGetTodoUseCase(repository)
	todo, err := useCases.FindById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errorMessage": err,
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}
