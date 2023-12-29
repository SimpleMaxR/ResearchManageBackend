package main

import (
	_ "ResearchManage/docs"
	"ResearchManage/internal/database"
	"ResearchManage/routes"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	routes.Init()
}
