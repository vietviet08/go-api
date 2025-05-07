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

	cfg.User = os.Getenv("DBUSER")
	if cfg.User == "" {
		cfg.User = appConfig.DB.User
	}

	cfg.Passwd = os.Getenv("DBPASS")
	if cfg.Passwd == "" {
		cfg.Passwd = appConfig.DB.Password
	}

	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", appConfig.DB.Host, appConfig.DB.Port)
	cfg.DBName = appConfig.DB.DBName

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
