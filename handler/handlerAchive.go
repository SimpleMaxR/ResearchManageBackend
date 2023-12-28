package handler

import (
	"ResearchManage/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Achievement struct {
		Name        string `json:"name" binding:"required"`
		Obtain      string `json:"obtain" binding:"required"`
		Contributor int32  `json:"contributor" binding:"required"`
		Project     int32  `json:"project" binding:"required"`
		Subtopoic   int32  `json:"subtopoic" binding:"required"`
		Type        int32  `json:"type" binding:"required"`
	}
)

func (apiCfg apiConfig) CreateAchievement(c *gin.Context) {

	// 解析参数
	if err := c.ShouldBindJSON(&Achievement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	achievement, err := apiCfg.DB.CreateAchievement(c.Request.Context(), database.CreateAchievementParams{
		Name:          Achievement.Name,
		Obtaineddate:  Achievement.Obtain,
		Contributorid: Achievement.Contributor,
		Baseproject:   Achievement.Project,
		Basesubtopic:  Achievement.Subtopoic,
		Type:          Achievement.Type,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":        "200",
		"msg":         "success",
		"achievement": achievement,
	})
}

func (apiCfg apiConfig) DeleteAchievement(c *gin.Context) {

	var achievementid int32
	// 解析参数
	if err := c.ShouldBindJSON(&achievementid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	err := apiCfg.DB.DeleteAchievement(c.Request.Context(), achievementid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
	})
}

func (apiCfg apiConfig) ListAchievementByProject(c *gin.Context) {

	var projectId int32

	// 解析参数
	if err := c.ShouldBindJSON(&projectId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	achievements, err := apiCfg.DB.ListAchievementByProject(c.Request.Context(), projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":         "200",
		"msg":          "success",
		"achievements": achievements,
	})
}

func (apiCfg apiConfig) ListAchievementBySubtopic(c *gin.Context) {

	var subtopicid int32
	// 解析参数
	if err := c.ShouldBindJSON(&subtopicid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	achievements, err := apiCfg.DB.ListAchievementBySubtopic(c.Request.Context(), subtopicid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":         "200",
		"msg":          "success",
		"achievements": achievements,
	})
}
