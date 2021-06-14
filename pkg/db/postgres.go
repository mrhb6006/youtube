package mysqlDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"os"
)

var connection *sql.DB

func setPostgresClient() {
	var err error
	URI := os.Getenv("POSTGRES")
	if URI == "" {
		URI = "127.0.0.1:5432"
	}
	USER := os.Getenv("POSTGRES_USER")
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB := os.Getenv("POSTGRES_DB")
	if USER == "" || PASSWORD == "" || DB == "" {
		zap.L().Error("postgres param required")
		os.Exit(1)
	}
	db, err := sql.Open("postgres", USER+":"+PASSWORD+"@tcp("+URI+")/"+DB)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	zap.L().Info("Successfully connected!")
}

func GetConnection() *sql.DB {
	if connection == nil {
		setPostgresClient()
	}
	return connection
}
