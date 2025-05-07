package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"

	"vietquoc/connect-db/config"
)

var DB *sql.DB

func InitDB() {
	cfg := mysql.NewConfig()

	appConfig := config.GetConfig()

	cfg.User = os.Getenv("DB_USER")
	if cfg.User == "" {
		cfg.User = appConfig.DB.User
	}

	cfg.Passwd = os.Getenv("DB_PASS")
	if cfg.Passwd == "" {
		cfg.Passwd = appConfig.DB.Password
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = appConfig.DB.Host
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = appConfig.DB.Port
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = appConfig.DB.DBName
	}

	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", dbHost, dbPort)
	cfg.DBName = dbName

	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database!")
}
