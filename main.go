package main

import (
	"ResearchManage/internal/database"
	"ResearchManage/routes"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	routes.Init()
}
