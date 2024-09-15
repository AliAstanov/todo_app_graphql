package postgres

import (
	"context"
	"example/graph/model"

	//model "example/graph/model"
	repoi "example/storage/repoI"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) repoi.UserRepoI {
	return &UserRepo{
		conn: conn,
	}
}

func (u *UserRepo) CreateUsers(ctx context.Context, user *model.User) (*model.User, error) {

	query := `
		INSERT INTO users(
			user_id,
			username, 
			password 
		) VALUES (
			$1, $2, $3
		)`

	_, err := u.conn.Exec(
		ctx, query,
		user.UserID,
		user.Username,
		user.Password,
	)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) GetUsers(ctx context.Context, limit, page string) ([]*model.User, error) {
	var users []*model.User

	// `limit` va `page` ni int turiga o'tkazish
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Invalid limit:", err)
		return nil, err
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Invalid page:", err)
		return nil, err
	}

	offset := (pageInt - 1) * limitInt

	query := `
		SELECT 
			user_id,	
			username, 
			password
		FROM 
			users
		LIMIT
			$1 
		OFFSET 
			$2
	`

	rows, err := u.conn.Query(ctx, query, limitInt, offset)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID,
			&user.Username,
			&user.Password); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error retrieving rows:", err)
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User

	query := `
		SELECT 
			user_id, 
			username, 
			password
		FROM 
			users
		WHERE 
			user_id = $1
		LIMIT 1
	`

	err := u.conn.QueryRow(ctx, query, userId).Scan(
		&user.UserID,
		&user.Username,
		&user.Password)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	query := `
		UPDATE 
			users
		SET 
			username = $1,
			password = $2
		WHERE 
			user_id = $3
	`

	_, err := u.conn.Exec(ctx, query,
		user.Username,
		user.Password,
		user.UserID)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) DeleteUserById(ctx context.Context, userID string) (*model.User, error) {
	var user model.User

	query := `
		DELETE FROM 
			users 
		WHERE 
			user_id = $1
		RETURNING 
			user_id, username
	`

	err := u.conn.QueryRow(ctx, query, userID).Scan(
		&user.UserID, 
		&user.Username)
	if err != nil {
		log.Println("Error deleting user:", err)
		return nil, err
	}

	return &user, nil
}


// CreateUsers(ctx context.Context, user *model.User) (*model.User, error)
// GetUsers(ctx context.Context, limit, page string) ([]*model.User, error)
// GetUserById(ctx context.Context, userId string) (*model.User, error)
// UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
// DeleteUserById(ctx context.Context, userID string) (*model.User, error)