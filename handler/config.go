package handler

import (
	"ResearchManage/internal/database"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

var apiCfg apiConfig

func InitDbConn() {
	dbURL := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	apiCfg = apiConfig{
		DB: database.New(conn),
	}
}

func GetApiCfg() *apiConfig {
	InitDbConn()
	return &apiCfg
}
