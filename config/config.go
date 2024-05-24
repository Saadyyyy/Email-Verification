package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func InitializeDatabase() (*sql.DB, error) {

	errE := godotenv.Load()
	if errE != nil {
		fmt.Println("Env can connect")
	}

	dbConfig := DatabaseConfig{}
	dbConfig.Host = os.Getenv("DBHOST")
	dbConfig.Port = os.Getenv("DBPORT")
	dbConfig.Username = os.Getenv("DBUSER")
	dbConfig.DBName = os.Getenv("DBNAME")
	dbConfig.Password = os.Getenv("DBPASS")
	// Konfigurasi koneksi database MySQL dengan GORM
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connect To Database")
	}

	return db, nil
}
