package postgres

import (
	"context"
	"example/graph/model"
	repoi "example/storage/repoI"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type TodoRepo struct {
	conn *pgx.Conn
}

func NewTodoRepo(conn *pgx.Conn) repoi.TodoRepoI {
	return &TodoRepo{conn: conn}
}

func (t *TodoRepo) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	query := `
		INSERT INTO
			todos(
				todo_id,
				user_id,
				task
			)VALUES(
				$1, $2, $3
			)`
	_, err := t.conn.Exec(
		ctx, query,
		todo.TodoID,
		todo.UserID,
		todo.Task)
	if err != nil {
		log.Println("error on CreateTodo:", err)
		return nil, err
	}

	return todo, nil
}
func (t *TodoRepo) GetTodos(ctx context.Context, limit, page string) ([]*model.Todo, error) {
	var todos []*model.Todo

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("invalid limit:", err)
		return nil, err
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println("invalid page")
		return nil, err
	}
	offset := (pageInt - 1) * limitInt

	query := `
		SELECT 
			todo_id,
			user_id,
			task
		FROM 
			todos
		LIMIT 
			$1
		OFFSET 
			$2
	`
	rows, err := t.conn.Query(ctx, query, limitInt, offset)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(
			&todo.TodoID,
			&todo.UserID,
			&todo.Task,
		); err != nil {
			log.Println("error on scaning todo:", err)
			return nil, err
		}

		todos = append(todos, &todo)
	}

	return todos, nil
}
func (t *TodoRepo) GetTodoById(ctx context.Context, todoId string) (*model.Todo, error) {
	var todo model.Todo

	query := `
		SELECT 
			todo_id,
			user_id,
			task
		FROM
			todos
		WHERE
			todo_id = $1
	`
	err := t.conn.QueryRow(ctx, query, todoId).Scan(
		&todo.TodoID,
		&todo.UserID,
		&todo.Task)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	return &todo, nil
}
func (t *TodoRepo) UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	query := `
		UPDATE
			todos
		SET
			task = $1
		WHERE
			todo_id = $2
		RETURNING
			todo_id,
			user_id,
			task		
	`
	var updateTodo model.Todo
	err := t.conn.QueryRow(ctx, query, todo.Task, todo.TodoID).Scan(
		&updateTodo.TodoID,
		&updateTodo.UserID,
		&updateTodo.Task)
	if err != nil {
		log.Println("error updating todo:", err)
		return nil, err
	}

	return &updateTodo, nil
}
func (t *TodoRepo) DeleteTodo(ctx context.Context, todoId string) (*model.Todo, error) {

	var todo model.Todo

	query := `
		DELETE FROM 
			todos
		WHERE
			todo_id = $1
		RETURNING
			todo_id, task
	`
	err := t.conn.QueryRow(ctx, query, todoId).Scan(
		&todo.TodoID,
		&todo.UserID,
		&todo.Task)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println("Todo not found with id:", todoId)
			return nil, fmt.Errorf("todo with id %s not found", todoId)
		}
		log.Println("error deleting todo:", err)
		return nil, err
	}
	return &todo, nil

}

