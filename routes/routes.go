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

	// 数据库检查接口
	router.GET("/healthzDatabase", apiCfg.HandlerHealthzDatabase)

	// Lab 相关接口
	router.GET("/listLabAll", apiCfg.HandlerListLabAll)
	router.GET("/listLabByLabID", apiCfg.HandlerListLabByLabID)
	router.POST("/createLab", apiCfg.HandlerCreateLab)
	router.DELETE("/deleteLab", apiCfg.HandlerDeleteLab)
	router.PUT("/updateLab", apiCfg.HandlerUpdateLab)
	//router.PUT("/updateLabOffice", apiCfg.HandlerUpdateLabOffice)

	// Researcher 相关接口
	router.GET("/listResearcherAll", apiCfg.HandlerListResearcherAll)
	router.POST("/createResearcher", apiCfg.HandlerCreateResearcher)
	router.DELETE("/deleteResearcher", apiCfg.HandlerDeleteResearcher)
	router.PUT("/updateResearcher", apiCfg.HandlerUpdateResearcher)

	// Secretary 相关接口
	router.GET("/listSecretaryAll", apiCfg.CreateSecretary)
	router.POST("/setSecretary", apiCfg.SetSecretary)
	router.DELETE("/deleteSecretary", apiCfg.DeleteSecretary)

	// Office 相关接口
	router.GET("/listOfficeAll", apiCfg.ListOfficeAll)
	router.POST("/createOffice", apiCfg.CreateOffice)
	router.DELETE("/deleteOffice", apiCfg.DeleteOffice)
	router.PUT("/updateOffice", apiCfg.UpdateOffice)

	// Project 相关接口
	router.GET("/listProjectAll", apiCfg.ListProjectAll)
	router.POST("/createProject", apiCfg.CreateProject)
	router.DELETE("/deleteProject", apiCfg.DeleteProject)
	router.PUT("/updateProject", apiCfg.UpdateProject)
	router.PUT("/linkProjectPartner", apiCfg.LinkProjectPartner)

	// Partner 相关接口
	router.POST("/createPartner", apiCfg.CreatePartner)
	router.GET("/listPartnerAll", apiCfg.ListPartners)
	router.GET("/getPartnerByProjectID", apiCfg.GetPartnerByProjectID)

	// Leader 相关接口
	router.POST("/createLeader", apiCfg.CreateLeader)
	router.GET("/getLeaderById", apiCfg.GetLeader)

	// QM 相关接口
	router.POST("/createQM", apiCfg.CreateQM)
	router.GET("/getQMByProjectID", apiCfg.GetQMByProjectID)
	router.PUT("setQMContact", apiCfg.SetQMContact)

	// Client 相关接口
	router.POST("/createClient", apiCfg.CreateClient)
	router.PUT("/updateClient", apiCfg.UpdateClient)
	router.PUT("/createClientContact", apiCfg.CreateClientContact)
	router.GET("/getClientByProjectID", apiCfg.GetClientByProjectID)

	// Achievement 相关接口
	router.POST("/createAchievement", apiCfg.CreateAchievement)
	router.DELETE("/deleteAchievement", apiCfg.DeleteAchievement)
	router.GET("/listAchievementByProject", apiCfg.ListAchievementByProject)
	router.GET("/listAchievementBySubtopic", apiCfg.ListAchievementBySubtopic)

	// 登陆
	router.POST("/login", apiCfg.HandlerAdminLogin)

	err := router.Run("localhost:" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
