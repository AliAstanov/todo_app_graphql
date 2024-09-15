package repoi

import (
	"context"
	"example/graph/model"
	//"github.com/google/uuid"
)

type UserRepoI interface {
	CreateUsers(ctx context.Context, user *model.User) (*model.User, error)
	GetUsers(ctx context.Context, limit, page string) ([]*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUserById(ctx context.Context, userID string) (*model.User, error)
}

type TodoRepoI interface {
	CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	GetTodos(ctx context.Context, limit, page string) ([]*model.Todo, error)
	GetTodoById(ctx context.Context, todoId string) (*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	DeleteTodo(ctx context.Context, todoId string) (*model.Todo, error)
}
