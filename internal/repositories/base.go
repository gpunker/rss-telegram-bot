package repositories

import (
	"database/sql"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetConnection() *sql.DB {
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = host + ":" + port
	cfg.DBName = os.Getenv("DBNAME")
	cfg.ParseTime = true

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}
