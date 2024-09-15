package db

import (
	"context"
	"fmt"
	"example/config"

	"github.com/jackc/pgx/v5"
)

var ctx = context.Background()

func ConnToDb(cfg config.Config) (*pgx.Conn, error) {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.PsqlConfig.User,
		cfg.PsqlConfig.Password,
		cfg.PsqlConfig.Host,
		cfg.PsqlConfig.Port,
		cfg.PsqlConfig.Database,
	)
	return pgx.Connect(ctx, dbUrl)
}

// package db

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// var DB *sql.DB

// func InitDB() {
// 	var err error
// 	connStr := "postgres://username:password@localhost:5432/dbname?sslmode=disable" // Replace with your DB info

// 	DB, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal("Failed to connect to the database:", err)
// 	}

// 	if err = DB.Ping(); err != nil {
// 		log.Fatal("Database connection is not alive:", err)
// 	}

// 	log.Println("Database connected successfully")
// }
