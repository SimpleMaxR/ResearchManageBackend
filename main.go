package main

import (
	"ResearchManage/handler"
	"ResearchManage/internal/database"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	// port := os.Getenv("PORT")

	// 配置 api
	apiCfg := handler.GetApiCfg()

	// 配置路由
	router := gin.Default()

	// 配置 CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Lab 相关接口
	router.GET("/listLabAll", apiCfg.HandlerListLabAll)
	router.POST("/createLab", apiCfg.HandlerCreateLab)
	router.GET("/healthzDatabase", apiCfg.HandlerHealthzDatabase)
	router.DELETE("/deleteLab", apiCfg.HandlerDeleteLab)
	router.PUT("/updateLab", apiCfg.HandlerUpdateLab)

	// Researcher 相关接口
	router.GET("/listResearcherAll", apiCfg.HandlerListResearcherAll)
	router.POST("/createResearcher", apiCfg.HandlerCreateResearcher)
	router.DELETE("/deleteResearcher", apiCfg.HandlerDeleteResearcher)
	router.PUT("/updateResearcher", apiCfg.HandlerUpdateResearcher)

	router.Run("localhost:" + os.Getenv("PORT"))

}
