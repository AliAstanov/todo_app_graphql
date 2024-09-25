package postgres

import (
	"example/config"
	"example/db"
	repoi "example/storage/repoI"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var pgRepo repoi.UserRepoI

func TestMain(m *testing.M) {
	cfg := config.Load()

	db, err := db.ConnToDb(cfg)
	if err != nil {
		log.Println("fail connection db on mainTest.go:", err)
		return
	}

	pgRepo = NewUserRepo(db)

	os.Exit(m.Run())
}
func CallDb() (*pgx.Conn, error) {
	cfg := config.Load()    // Yana konfiguratsiyani yuklash
	return db.ConnToDb(cfg) // Ma'lumotlar bazasiga ulanish
}
