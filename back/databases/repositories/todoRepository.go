package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tomo-micco/TodoWithGo/databases/entities"
)

type TodoRepositoryInterface interface {
	GetAll(c context.Context) []entities.Todo
	FindById(c context.Context, id int64) entities.Todo
}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepositoryInterface {
	return &TodoRepository{db}
}

/*
* 全件取得
 */
func (repository *TodoRepository) GetAll(c context.Context) []entities.Todo {
	sql := " SELECT * FROM todos; "
	rows, err := repository.db.QueryContext(c, sql)
	if err != nil {
		fmt.Printf("error occurred in execute query: %v", err)
		return nil
	}

	defer rows.Close()
	var todos []entities.Todo
	for rows.Next() {
		var todo entities.Todo
		err := rows.Scan(&todo.Id, &todo.Content, &todo.IsComplete)
		if err != nil {
			fmt.Printf("error occurred in scan rows: %v", err)
			return nil
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("error occurred in rows: %v", err)
		return nil
	}

	return todos
}

/*
* idに該当するTodoを取得
 */
func (repository *TodoRepository) FindById(c context.Context, id int64) entities.Todo {
	sql := "SELECT * FROM todos WHERE id = ?"

	var todo entities.Todo
	err := repository.db.QueryRowContext(c, sql, id).Scan(&todo.Id, &todo.Content, &todo.IsComplete)
	if err != nil {
		fmt.Printf("error occurred in execute query: %v", err)
		return todo
	}

	return todo
}
