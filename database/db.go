package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DB_HOST")
	cfg.DBName = os.Getenv("DB_NAME")
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("erro ao conectar")
	}
	if err := db.Ping(); err != nil {
		log.Fatal("erro ao testar conection")
	}
	fmt.Println("Conectado com sucesso ao MySQL!")
	return db
}
