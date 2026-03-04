package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func EstablishConnection() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file: ", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		err := db.Ping()
		if err != nil {
			log.Print("Cannot reach db: ", err)
		} else {
			return db
		}
		time.Sleep(time.Second * 5)
	}

	log.Fatal("Error cannot reach db.")
	return db
}
