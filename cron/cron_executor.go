package cron

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/joho/godotenv"
	"github.com/muhammadisa/go-cron-service/cron/cronmaker"
	"log"
	"os"
	"strings"
)

func loadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println("Environment loaded")
}

func databaseConnect() *sql.DB {
	driverAndStrConn := fmt.Sprintf("mysql~%s", fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	connectionStr := strings.Split(driverAndStrConn, "~")
	db, err := sql.Open(connectionStr[0], connectionStr[1])
	if err != nil {
		log.Println(err)
		os.Exit(-1)
		return nil
	}
	log.Println("Connected to database")
	return db
}

func createDBRSession(db *sql.DB) *dbr.Session {
	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}
	conn.SetMaxOpenConns(10)
	session := conn.NewSession(nil)
	_, err := session.Begin()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
		return nil
	}
	return session
}

// Run start run cron scheduler
func Run() {
	// Load environment
	loadEnvironment()
	// Connecting to database
	db := databaseConnect()
	// Creating dbr session
	session := createDBRSession(db)
	// Initialize crons
	cronmaker.
		InitCrons(session).
		StartCronJobs()
}
