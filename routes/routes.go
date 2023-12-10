package routes

import (
	"ResearchManage/handler"
	"ResearchManage/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load(".env")
	// port := os.Getenv("PORT")

	// 配置 api
	apiCfg := handler.GetApiCfg()

	// 配置路由
	router := gin.Default()

	// 配置 CORS
	router.Use(middleware.Cors())

	// 配置错误处理
	router.Use(gin.Recovery())

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

	// 系统管理员
	Admin := router.Group("api/v1/admin")
	Admin.POST("/login", apiCfg.HandlerAdminLogin)

	router.Run("localhost:" + os.Getenv("PORT"))
}
