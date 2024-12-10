package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tomo-micco/TodoWithGo/databases/entities"
)

type TodoRepositoryInterface interface {
	GetAll(c context.Context) ([]entities.Todo, error)
	FindById(c context.Context, id int64) (entities.Todo, error)
	Create(c context.Context, todo entities.Todo) (int64, error)
	Update(c context.Context, todo entities.Todo) (int64, error)
	Delete(c context.Context, id int64) (int64, error)
}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepositoryInterface {
	return &TodoRepository{db}
}

/*
全件取得
*/
func (repository *TodoRepository) GetAll(c context.Context) ([]entities.Todo, error) {
	sql := " SELECT * FROM todos; "
	rows, err := repository.db.QueryContext(c, sql)
	if err != nil {
		fmt.Printf("error occurred in execute query: %v", err)
		return nil, err
	}

	defer rows.Close()

	var todos []entities.Todo
	for rows.Next() {
		var todo entities.Todo
		err := rows.Scan(&todo.Id, &todo.Content, &todo.IsComplete)
		if err != nil {
			fmt.Printf("error occurred in scan rows: %v", err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("error occurred in rows: %v", err)
		return nil, err
	}

	return todos, nil
}

/*
idに該当するTodoを取得
*/
func (repository *TodoRepository) FindById(c context.Context, id int64) (entities.Todo, error) {
	sql := "SELECT * FROM todos WHERE id = ?"

	var todo entities.Todo
	err := repository.db.QueryRowContext(c, sql, id).Scan(&todo.Id, &todo.Content, &todo.IsComplete)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

/*
Todoの作成
*/
func (repository *TodoRepository) Create(c context.Context, todo entities.Todo) (int64, error) {
	sql := `
		INSERT INTO
			todos (
				todo,
				is_complete
			) VALUES (
				?,
				?
			); `

	tx, err := repository.db.BeginTx(c, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // パニック時にはロールバックする
			panic(p)
		} else if err != nil {
			_ = tx.Rollback() // エラー発生時もロールバック
		} else {
			err = tx.Commit() // 正常終了時はコミット
		}
	}()

	result, err := tx.ExecContext(c, sql, todo.Content, todo.IsComplete)
	if err != nil {
		return 0, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedId, nil
}

/*
更新処理
*/
func (repository *TodoRepository) Update(c context.Context, todo entities.Todo) (int64, error) {
	sql := `
		UPDATE 
			todos 
		SET 
			todo = ?,
			is_complete = ?
		WHERE 
			id = ?
	`

	tx, err := repository.db.BeginTx(c, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // パニック時はロールバック
			panic(c)
		} else if err != nil {
			_ = tx.Rollback() // エラー発生時もロールバック
		} else {
			err = tx.Commit() // 正常終了時はコミット
		}
	}()

	result, err := tx.ExecContext(c, sql, todo.Content, todo.IsComplete, todo.Id)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

/*
削除処理
*/
func (repository *TodoRepository) Delete(c context.Context, id int64) (int64, error) {
	sql := "DELETE FROM todos WHERE id = ?;"

	tx, err := repository.db.BeginTx(c, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // パニック時はロールバック
			panic(c)
		} else if err != nil {
			_ = tx.Rollback() // エラー発生時もロールバック
		} else {
			err = tx.Commit() // 正常終了時はコミット
		}
	}()

	result, err := tx.ExecContext(c, sql, id)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
