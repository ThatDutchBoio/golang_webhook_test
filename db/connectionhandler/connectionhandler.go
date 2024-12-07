package connectionhandler

import (
	"database/sql"
	"fmt"
	"log"
	"main/apikeys"
	"main/utils/hashing"
	"os"

	"github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	hashing.SHA256("testing")
	key := apikeys.GenerateApiKey()
	fmt.Println("webhook id:")
	fmt.Println(key)
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "webhooks",
		AllowNativePasswords: true,
	}

	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)

	}
	fmt.Println("Connected!")

	return db
}
