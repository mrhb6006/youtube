package mysqlDB

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os"
	"strconv"
)

var connection *sql.DB

func setPostgresClient() {
	var err error
	HOST := os.Getenv("POSTGRES")
	if HOST == "" {
		HOST = "127.0.0.1"
	}
	P := os.Getenv("PORT")
	PORT, _ := strconv.Atoi(P)
	USER := os.Getenv("POSTGRES_USER")
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB := os.Getenv("POSTGRES_DB")
	if USER == "" || PASSWORD == "" || DB == "" {
		zap.L().Error("postgres param required")
		os.Exit(1)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DB)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	zap.L().Info("Successfully connected!")
}

func GetPGConnection() *sql.DB {
	if connection == nil {
		setPostgresClient()
	}
	return connection
}
